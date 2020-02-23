package server

import (
	"context"
	"fmt"
	"github.com/angenalZZZ/Go/go-gin-api/server/config"
	"github.com/angenalZZZ/Go/go-gin-api/server/route"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run() {
	conf := config.AppConfig

	gin.SetMode(conf.AppMode)
	engine := gin.New()

	// 性能分析 - 正式环境不要使用！！！
	pprof.Register(engine)

	// 设置路由
	route.SetupRouter(engine)

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", conf.Server.ListenAddr, conf.Server.Port),
		Handler:      engine,
		ReadTimeout:  time.Duration(conf.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(conf.Server.WriteTimeout) * time.Second,
	}

	fmt.Println("|-----------------------------------|")
	fmt.Println("|            go-gin-api             |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("|  Go Http Server Start Successful  |")
	fmt.Println("|    Server:" + server.Addr + "     Pid:" + fmt.Sprintf("%d", os.Getpid()) + "        |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("")

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	log.Println("Get Signal:", sig)
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
