package go_scheduler

import (
	"angenalZZZ/go-program/api-config"
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"log"
	"os"
	"strconv"
)

// Make a Worker Pool
var WorkerPool *work.WorkerPool

// Make a redis pool
var cliPool *redis.Pool
var cliOpt redis.DialOption
var cliAddr string

type Context struct {
	CustomerID string
}

func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	log.Printf(" [%s] job:Starting...\n ", job.Name)
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
	// Extract arguments:
	address, subject := job.ArgString("address"), job.ArgString("subject")
	if err := job.ArgError(); err != nil {
		return err
	}

	// SendEmailTo(addr, subject)
	log.Printf(" [%s] job:SendEmail... \n  address: %s\n  subject: %s \n", job.Name, address, subject)

	return nil
}

func (c *Context) Export(job *work.Job) error {
	return nil
}

// 初始化配置
func init() {
	// config
	api_config.Check("REDIS_ADDR")
	api_config.Check("REDIS_PWD")
	api_config.Check("REDIS_DB")
	cliAddr = os.Getenv("REDIS_ADDR")
	i, e := strconv.Atoi(os.Getenv("REDIS_DB"))
	if e != nil {
		i = 0
	}
	// client
	cliOpt = redis.DialClientName("redis-cli")
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
}
func initCliPool() {
	if cliPool != nil {
		return
	}

	// managed Pool
	cliPool = &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", cliAddr, cliOpt)
		},
	}
}
func initWorkerPool() {
	if WorkerPool != nil {
		return
	}

	initCliPool()

	c := Context{}
	WorkerPool = work.NewWorkerPool(c, 10, "app_namespace", cliPool)

	// Add middleware that will be executed for each job
	WorkerPool.Middleware((*Context).Log)
	WorkerPool.Middleware((*Context).FindCustomer)

	// Map the name of jobs to handler functions
	WorkerPool.Job("send_email", (*Context).SendEmail)

	// Customize options:
	WorkerPool.JobWithOptions("export", work.JobOptions{Priority: 10, MaxFails: 1}, (*Context).Export)

	// Start processing jobs
	WorkerPool.Start()
}
