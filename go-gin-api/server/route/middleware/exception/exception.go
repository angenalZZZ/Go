package exception

import (
	"fmt"
	"github.com/angenalZZZ/Go/go-gin-api/server/config"
	"github.com/angenalZZZ/Go/go-gin-api/server/util/ctx"
	"github.com/angenalZZZ/gofunc/f"
	"github.com/gin-gonic/gin"
	"runtime/debug"
	"strings"
	"time"
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
				body = strings.ReplaceAll(body, "{RequestTime}", time.Now().Format("2006/01/02 15:04:05"))
				body = strings.ReplaceAll(body, "{RequestURL}", c.Request.Method+"  "+c.Request.Host+c.Request.RequestURI)
				body = strings.ReplaceAll(body, "{RequestUA}", c.Request.UserAgent())
				body = strings.ReplaceAll(body, "{RequestIP}", c.ClientIP())
				body = strings.ReplaceAll(body, "{DebugStack}", DebugStack)

				// 执行发邮件
				options := &f.MailOptions{
					MailSMTP: f.MailSMTP{
						Port: conf.NotifyUser.Smtp.Port,
						Host: conf.NotifyUser.Smtp.Host,
						User: conf.NotifyUser.Smtp.User,
						Pass: conf.NotifyUser.Smtp.Pass,
					},
					MailMessage: f.MailMessage{
						Recipient: []string{conf.NotifyUser.ErrorNotifyUser},
						Subject:   subject,
						Body:      body,
					},
				}
				_ = f.MailSend(options)

				ctx.Wrap(c).Error("系统异常，请联系管理员！", nil)
			}
		}()
		c.Next()
	}
}
