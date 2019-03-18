package go_scheduler

import (
	"angenalZZZ/go-program/api-config"
	"fmt"
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"log"
	"os"
	"strconv"
	"time"
)

// Make a Worker Pool
var WorkerPool *work.WorkerPool
var workerConfig WorkerConfig

// Worker Config
type WorkerConfig struct {
	concurrency   uint
	namespace     string // eg, "myapp-work"
	pool          *redis.Pool
	sleepBackoffs []int64
}

// Worker Context
type Context struct {
	CustomerID string
}

func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	//log.Printf(" [%s] job:Starting...\n ", job.Name)
	return next()
}

func (c *Context) FindCustomer(job *work.Job, next work.NextMiddlewareFunc) error {
	// If there's a customer_id param, set it in the context for future middleware and handlers to use.
	if _, ok := job.Args["customer_id"]; ok {
		c.CustomerID = job.ArgString("customer_id")
		if err := job.ArgError(); err != nil {
			return err
		}
	}
	return next()
}

func (c *Context) SendEmail(job *work.Job) error {
	// Manual Checkin with args, Enqueue
	log.Printf(" job: %+v", job)

	if !job.Unique {

		// Extract arguments
		address, subject := job.ArgString("address"), job.ArgString("subject")
		if err := job.ArgError(); err != nil {
			return err
		}

		// SendEmailTo(addr, subject)
		log.Printf(" [%s] job:SendEmail[Manual] OK \n  address: %s\n  subject: %s \n", job.Name, address, subject)

		return nil
	}

	// Auto Checkin with no args, Worker Pool
	log.Printf(" [%s] job:SendEmail[Auto] OK \n", job.Name)

	return nil
}

func (c *Context) Export(job *work.Job) error {
	job.Checkin(fmt.Sprintf("CustomerID=%s : export", c.CustomerID))
	return nil
}

// 初始化配置
func init() {
	// config
	api_config.Check("REDIS_ADDR")
	api_config.Check("REDIS_PWD")
	api_config.Check("REDIS_DB")
	cliAddr := os.Getenv("REDIS_ADDR")
	i, e := strconv.Atoi(os.Getenv("REDIS_DB"))
	if e != nil {
		i = 0
	}
	// client
	cliOpt := redis.DialClientName("redis-cli")
	cliOpt = redis.DialUseTLS(false)
	// password
	password := os.Getenv("REDIS_PWD")
	if len(password) > 0 {
		cliOpt = redis.DialPassword(password)
	}
	// db number
	if i > 0 && i < 16 {
		cliOpt = redis.DialDatabase(i)
	}
	// managed Pool
	cliPool := &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", cliAddr, cliOpt)
		},
	}
	// Worker Config
	workerConfig = WorkerConfig{
		concurrency:   1,
		namespace:     "work",
		pool:          cliPool,
		sleepBackoffs: []int64{},
	}
}

func initWorkerPool() {
	if WorkerPool != nil {
		return
	}

	c := Context{CustomerID: fmt.Sprintf("C%d", time.Now().Year())}
	workerPoolOpts := work.WorkerPoolOptions{}
	WorkerPool = work.NewWorkerPoolWithOptions(c, workerConfig.concurrency, workerConfig.namespace, workerConfig.pool, workerPoolOpts)

	// 任务之前 Add middleware that will be executed for each job
	WorkerPool.Middleware((*Context).Log)
	WorkerPool.Middleware((*Context).FindCustomer)

	// 添加任务/发送邮件 Map the name of jobs to handler functions
	WorkerPool.Job("send_email", (*Context).SendEmail)

	// 添加任务/导出数据 Customize options
	WorkerPool.JobWithOptions("export", work.JobOptions{Priority: 10, MaxFails: 1, MaxConcurrency: 1}, (*Context).Export)
}

// 计划任务 close
func ShutdownWorker() {
	if WorkerPool != nil {
		//log.Println("计划任务 Scheduler Worker closing..")
		WorkerPool.Stop()
	}
}

// 测试
func TestWorker() {
	initWorkerPool()

	/**
	自动化任务
	*/
	// 每隔2秒 https://godoc.org/github.com/robfig/cron
	WorkerPool.PeriodicallyEnqueue("@every 2s", "send_email")
	// 每小时 @hourly
	WorkerPool.PeriodicallyEnqueue("0 0 * * * *", "export")
	// Start processing jobs
	WorkerPool.Start()

	/**
	手动任务
	*/
	enqueuer := work.NewEnqueuer(workerConfig.namespace, workerConfig.pool)
	// 添加任务/发送邮件
	q := work.Q{"address": "123456.qq.com", "subject": "some message"}
	job, e := enqueuer.EnqueueIn("send_email", 10, q)
	if e != nil {
		log.Printf(" [%s] job:SendEmail[Manual]... \n  address: %s\n  subject: %s \n %v \n", job.Name, q["address"], q["subject"], e)
	} else {
		log.Printf(" [%s] job:SendEmail[Manual]... \n  address: %s\n  subject: %s \n", job.Name, q["address"], q["subject"])
	}
}
