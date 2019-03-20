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
var WorkerPool *work.WorkerPool

// 任务执行器的配置
var workerPoolOpts work.WorkerPoolOptions

// 任务的配置
var workerConfig WorkerConfig

// 任务的上下文对象
var workerCtx Context

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

	// 任务的配置
	workerConfig = WorkerConfig{
		concurrency:     1000, // 并发数,默认0不限制
		namespace:       "work",
		pool:            cliPool,
		sleepBackoffs:   []int64{},
		log:             true, // 开启当前上下文的日志输出Context.log
		drainOnShutdown: true, // 计划任务 close 时, 清理队列
	}

	// 任务执行器的配置
	workerPoolOpts = work.WorkerPoolOptions{}

	// 任务的上下文对象
	workerCtx = Context{customer: fmt.Sprintf("C%d", time.Now().Year())}
}

// 初始化 Worker Pool
func initWorkerPool() {
	if WorkerPool != nil {
		return
	}

	// 任务的上下文对象
	WorkerPool = work.NewWorkerPoolWithOptions(workerCtx, workerConfig.concurrency, workerConfig.namespace, workerConfig.pool, workerPoolOpts)

	// 任务执行时,通用过滤器 Add middleware that will be executed for each job
	WorkerPool.Middleware(workerCtx.Log)               // 日志跟踪
	WorkerPool.Middleware(workerCtx.FindCtxUseJobAttr) // 通过查找参数补充任务的上下文信息
}

// 初始化 Worker Jobs
func initWorkerJobs() {
	if WorkerPool == nil {
		initWorkerPool()
	}

	// 添加任务/发送邮件 Map the name of jobs to handler functions
	WorkerPool.Job("send_email", workerCtx.SendEmail)

	// 添加并发任务/导出数据 Customize options
	WorkerPool.JobWithOptions("export", work.JobOptions{Priority: 10, MaxFails: 0, MaxConcurrency: workerConfig.concurrency}, workerCtx.Export)
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
	定期执行任务 Worker Jobs
	*/
	if workerCtx.manual == false {
		// 每1秒  参考(spec: 秒/分/时/日/月/星期) https://godoc.org/github.com/robfig/cron
		WorkerPool.PeriodicallyEnqueue("*/1 * * * * *", "send_email")
		// 每1小时 @hourly
		WorkerPool.PeriodicallyEnqueue("0 0 * * * *", "export")
	}

	/**
	手动执行任务 Worker Jobs
	*/
	if workerCtx.manual == true {
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

// 日志跟踪
func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	//c.log(job, "Log")
	return next()
}

// 通过查找参数补充任务的上下文信息
func (c *Context) FindCtxUseJobAttr(job *work.Job, next work.NextMiddlewareFunc) error {
	if e := c.Check(job); e != nil {
		return e
	}
	//c.log(job, "FindCustomer")
	return next()
}

// 任务：发送邮件
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
		c.Checkin(job)
		return nil
	}

	// Auto Checkin with no args, Worker Pool
	log.Printf(" job[%s]:Work End one time. \n  %+v \n ", job.Name, c)
	c.Checkin(job)
	return nil
}

// 任务：导出功能
func (c *Context) Export(job *work.Job) error {
	c.Checkin(job)
	return nil
}

// 任务配置项 Worker Config
type WorkerConfig struct {
	concurrency     uint
	namespace       string // eg, "myapp-work"
	pool            *redis.Pool
	sleepBackoffs   []int64
	log             bool // 开启当前上下文的日志输出Context.log
	drainOnShutdown bool // 计划任务 close 时, 清理队列
}

// 任务上下文对象 Worker Context
type Context struct {
	customer string // 客户
	manual   bool   // 是否为 手动执行任务 run method: cron or manual
}

func (c *Context) log(job *work.Job, method string) {
	if workerConfig.log {
		log.Printf(" job[%s]:%s \n  %+v \n  %+v \n", job.Name, method, c, job)
	}
}

func (c *Context) Check(job *work.Job) error {
	// If there's a customer param, set it in the context for future middleware and handlers to use.
	if _, ok := job.Args["customer"]; ok {
		c.customer = job.ArgString("customer")
		if err := job.ArgError(); err != nil {
			return err
		}
	}
	// run method: cron or manual
	if _, ok := job.Args["cron"]; ok {
		c.manual = false
	}
	if _, ok := job.Args["manual"]; ok {
		c.manual = true
	}
	return nil
}

func (c *Context) Checkin(job *work.Job) {
	job.Checkin(fmt.Sprintf("job[%s]%s:#%s", c.customer, job.Name, job.ID))
}

func UniqueWorkQ() work.Q {
	id := uuid.Must(uuid.NewV4())
	return work.Q{"uuid": id.String()} // object_id_
}
