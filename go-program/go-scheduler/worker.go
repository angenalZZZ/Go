package go_scheduler

import (
	"angenalZZZ/go-program/api-config"
	"fmt"
	"github.com/gocraft/work"
	"github.com/gofrs/uuid"
	"github.com/gomodule/redigo/redis"
	"log"
	"os"
	"strconv"
	"time"
)

// 计划任务执行器
// Make a Worker Pool
var WorkerPool *work.WorkerPool
var workerConfig WorkerConfig

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
		concurrency:     1,
		namespace:       "work",
		pool:            cliPool,
		sleepBackoffs:   []int64{},
		log:             true, // 开启当前上下文的日志输出Context.log
		manual:          true, // 是否为 手动执行任务
		drainOnShutdown: true, // 计划任务 close 时, 清理队列
	}
}

// 初始化 Worker Pool
func initWorkerPool() {
	if WorkerPool != nil {
		return
	}

	// 自动化任务上下文对象
	c := Context{customer: fmt.Sprintf("C%d", time.Now().Year())}
	workerPoolOpts := work.WorkerPoolOptions{}
	WorkerPool = work.NewWorkerPoolWithOptions(c, workerConfig.concurrency, workerConfig.namespace, workerConfig.pool, workerPoolOpts)

	// 任务执行时通用过滤器 Add middleware that will be executed for each job
	WorkerPool.Middleware(c.Log)
	WorkerPool.Middleware(c.FindCustomer)
}

// 初始化 Worker Jobs
func initWorkerJobs() {
	if WorkerPool == nil {
		initWorkerPool()
	}

	// 添加任务/发送邮件 Map the name of jobs to handler functions
	WorkerPool.Job("send_email", (*Context).SendEmail)

	// 添加任务/导出数据 Customize options
	WorkerPool.JobWithOptions("export", work.JobOptions{Priority: 10, MaxFails: 1, MaxConcurrency: 1}, (*Context).Export)
}

// 计划任务 close
func ShutdownWorker() {
	if WorkerPool != nil {
		// 计划任务 close 时, 清理队列
		if workerConfig.drainOnShutdown {
			WorkerPool.Drain()
		}
		//log.Println("计划任务 Scheduler Worker closing..")
		WorkerPool.Stop()
	}
}

// 测试
func TestWorker() {
	/**
	初始化 Worker Jobs
	*/
	initWorkerJobs()

	// 开始按计划执行 Start processing jobs
	defer WorkerPool.Start()

	/**
	计划执行任务 Worker Jobs
	*/
	if workerConfig.manual == false {
		// 每1秒  参考(spec: 秒/分/时/日/月/星期) https://godoc.org/github.com/robfig/cron
		WorkerPool.PeriodicallyEnqueue("*/1 * * * * *", "send_email")
		// 每1小时 @hourly
		WorkerPool.PeriodicallyEnqueue("0 0 * * * *", "export")
	}

	/**
	手动执行任务 Worker Jobs
	*/
	if workerConfig.manual == true {
		enqueuer := work.NewEnqueuer(workerConfig.namespace, workerConfig.pool)
		// 添加任务/发送邮件
		q, id := work.Q{"address": "123456.qq.com", "subject": "some message"}, UniqueWorkQ()
		job, e := enqueuer.EnqueueUniqueByKey("send_email", q, id)
		if e != nil {
			log.Printf(" job[%s]:[Manual]... Err \n  address: %s\n  subject: %s \n %v \n", job.Name, q["address"], q["subject"], e)
		} else {
			log.Printf(" job[%s]:[Manual]... OK \n  address: %s\n  subject: %s \n", job.Name, q["address"], q["subject"])
		}
	}
}

// 任务配置项 Worker Config
type WorkerConfig struct {
	concurrency     uint
	namespace       string // eg, "myapp-work"
	pool            *redis.Pool
	sleepBackoffs   []int64
	log             bool // 开启当前上下文的日志输出Context.log
	manual          bool // 是否为 手动执行任务
	drainOnShutdown bool // 计划任务 close 时, 清理队列
}

// 任务上下文对象 Worker Context
type Context struct {
	customer string
	manual   bool // 是否为 手动执行任务
}

func (c *Context) log(job *work.Job, method string) {
	if workerConfig.log {
		log.Printf(" job[%s]:%s \n  %+v \n  %+v \n", job.Name, method, c, job)
	}
}

func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	//c.log(job, "Log")
	return next()
}

func (c *Context) FindCustomer(job *work.Job, next work.NextMiddlewareFunc) error {
	if e := c.Check(job); e != nil {
		return e
	}
	c.log(job, "FindCustomer")
	return next()
}

func (c *Context) SendEmail(job *work.Job) error {
	if e := c.Check(job); e != nil {
		return e
	}
	c.log(job, "SendEmail")

	// Manual Checkin with args and Enqueue
	if c.manual == true {
		log.Printf(" job[%s]:Work Start manual. \n", job.Name)

		// Extract arguments
		address, subject := job.ArgString("address"), job.ArgString("subject")
		if err := job.ArgError(); err != nil {
			log.Printf(" job[%s]:Work Error manual. \n  %v \n", job.Name, err)
			return err
		}

		// SendEmailTo(addr, subject)
		log.Printf(" job[%s]:Work End manual. \n  %+v \n  address: %s\n  subject: %s \n", job.Name, c, address, subject)

		return nil
	}

	// Auto Checkin with no args, Worker Pool
	log.Printf(" job[%s]:Work End one time. \n  %+v \n ", job.Name, c)

	return nil
}

func (c *Context) Export(job *work.Job) error {
	job.Checkin(fmt.Sprintf("customer=%s:export", c.customer))
	return nil
}

func (c *Context) Check(job *work.Job) error {
	// If there's a customer param, set it in the context for future middleware and handlers to use.
	c.manual = workerConfig.manual
	if _, ok := job.Args["customer"]; ok {
		c.customer = job.ArgString("customer")
		if err := job.ArgError(); err != nil {
			return err
		}
	} else {
		if job.Args == nil {
			job.Args = work.Q{"customer": c.customer}
		} else {
			job.Args["customer"] = c.customer
		}
	}
	return nil
}

func UniqueWorkQ() work.Q {
	id := uuid.Must(uuid.NewV4())
	return work.Q{"uuid": id.String()} // object_id_
}
