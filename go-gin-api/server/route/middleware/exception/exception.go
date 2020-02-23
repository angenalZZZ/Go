package exception

import (
	"fmt"
	"github.com/angenalZZZ/Go/go-gin-api/server/config"
	"github.com/angenalZZZ/Go/go-gin-api/server/util/response"
	"github.com/gin-gonic/gin"
	"github.com/xinliangnote/go-util/mail"
	"github.com/xinliangnote/go-util/time"
	"runtime/debug"
	"strings"
)

func SetUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				conf := config.AppConfig

				DebugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					DebugStack += v + "<br>"
				}

				subject := fmt.Sprintf("【重要错误】%s 项目出错了！", conf.AppName)

				body := strings.ReplaceAll(MailTemplate, "{ErrorMsg}", fmt.Sprintf("%s", err))
				body = strings.ReplaceAll(body, "{RequestTime}", time.GetCurrentDate())
				body = strings.ReplaceAll(body, "{RequestURL}", c.Request.Method+"  "+c.Request.Host+c.Request.RequestURI)
				body = strings.ReplaceAll(body, "{RequestUA}", c.Request.UserAgent())
				body = strings.ReplaceAll(body, "{RequestIP}", c.ClientIP())
				body = strings.ReplaceAll(body, "{DebugStack}", DebugStack)

				options := &mail.Options{
					MailHost: conf.NotifyUser.Smtp.Host,
					MailPort: conf.NotifyUser.Smtp.Port,
					MailUser: conf.NotifyUser.Smtp.User,
					MailPass: conf.NotifyUser.Smtp.Pass,
					MailTo:   conf.NotifyUser.ErrorNotifyUser,
					Subject:  subject,
					Body:     body,
				}
				_ = mail.Send(options)

				utilGin := response.Gin{Ctx: c}
				utilGin.Response(500, "系统异常，请联系管理员！", nil)
			}
		}()
		c.Next()
	}
}
