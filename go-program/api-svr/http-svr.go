package api_svr

import (
	"context"
	"log"
	"net"
	"net/http"

	api_config "github.com/angenalZZZ/Go/go-program/api-config"
	"github.com/angenalZZZ/Go/go-program/api-svr/authtoken"
	"github.com/angenalZZZ/Go/go-program/api-svr/img"
	gormMysql "github.com/angenalZZZ/Go/go-program/api-svr/orm/gorm/mysql"
	gormSqlite "github.com/angenalZZZ/Go/go-program/api-svr/orm/gorm/sqlite"
	sqlxSqlite "github.com/angenalZZZ/Go/go-program/api-svr/orm/sqlx/sqlite"
)

/**
后台服务 http: Server
*/
var httpSvr *http.Server

func initHttpSvr() {
	httpSvr = &http.Server{Addr: api_config.Config.HttpSvr.Addr}
}

// 后台运行 http Serve Run
func TestHttpSvrRun() {
	initHttpSvr()

	// 静态资源访问 html,css,js...
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// 服务处理：验证码
	http.HandleFunc("/api/captcha/get", img.CaptchaGenerateHandler)
	http.HandleFunc("/api/captcha/verify", img.CaptchaVerifyHandle)

	// 账号信息认证：AUTH JWT
	http.HandleFunc("/token/jwt", authtoken.JwtTokenGenerateHandler)
	http.HandleFunc("/token/jwt/verify", authtoken.JwtVerifyValidateHandler)
	http.HandleFunc("/token/jwt/sign", authtoken.JsonSignGenerateHandler)
	http.HandleFunc("/token/jwt/sign/verify", authtoken.JsonSignValidateHandler)

	// 数据库 gorm
	http.HandleFunc("/gorm/mysql/test", gormMysql.FooTestHandler)
	http.HandleFunc("/gorm/sqlite/test", gormSqlite.FooTestHandler)
	// 数据库 sqlx
	http.HandleFunc("/sqlx/sqlite/test", sqlxSqlite.FooTestHandler)

	// 服务处理
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	// 跟踪请求
	//	log.Printf("后台服务 http: %s\n", r.URL)
	//	// 处理请求
	//	_, e := fmt.Fprintf(w, " %v %+v \n", time.Now(), r.URL)
	//	// 跟踪异常
	//	if e != nil {
	//		log.Println(e)
	//	}
	//})

	// 开始服务
	//log.Fatal(http.ListenAndServe(httpSvr.Addr, nil)) // a simple way
	l, e := net.Listen("tcp4", httpSvr.Addr)
	if e == nil {
		println()
		log.Printf("后台服务 http: Server starting.. Addr: %s\n\n", httpSvr.Addr)
		if e = httpSvr.Serve(l); e != nil {
			log.Fatal(e) // 中断程序时输出
		}
	} else {
		log.Fatal(e) // 中断程序时输出
	}
}

// 后台运行 http Serve Shutdown
func HttpSvrShutdown() {
	if httpSvr != nil {
		log.Println("后台服务 http: Server stopping..") // Go ^1.8
		if e := httpSvr.Shutdown(context.Background()); e != nil {
			log.Fatal(e) // 中断程序时输出
		}
	}
}
