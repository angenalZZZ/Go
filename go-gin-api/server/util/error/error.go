package error

import (
	"errors"
	"fmt"
	"github.com/angenalZZZ/Go/go-gin-api/server/config"
	"github.com/angenalZZZ/Go/go-gin-api/server/route/middleware/exception"
	"github.com/angenalZZZ/gofunc/log"
	"github.com/angenalZZZ/gofunc/log/lager"
	"github.com/xinliangnote/go-util/mail"
	timeUtil "github.com/xinliangnote/go-util/time"
	"runtime/debug"
	"strings"
)

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func ErrorNew(text string) error {
	alarm("INFO", text)
	return &errorString{text}
}

// 发邮件
func ErrorMail(text string) error {
	alarm("MAIL", text)
	return &errorString{text}
}

// 发短信
func ErrorSms(text string) error {
	alarm("SMS", text)
	return &errorString{text}
}

// 发微信
func ErrorWeChat(text string) error {
	alarm("WX", text)
	return &errorString{text}
}

// 告警方法
func alarm(level string, str string) {
	conf := config.AppConfig

	debugStack := ""
	for _, v := range strings.Split(string(debug.Stack()), "\n") {
		debugStack += v + "<br>"
	}

	subject := fmt.Sprintf("【系统告警】%s 项目出错了！", conf.AppName)

	body := strings.ReplaceAll(exception.MailTemplate, "{ErrorMsg}", fmt.Sprintf("%s", str))
	body = strings.ReplaceAll(body, "{RequestTime}", timeUtil.GetCurrentDate())
	body = strings.ReplaceAll(body, "{RequestURL}", "--")
	body = strings.ReplaceAll(body, "{RequestUA}", "--")
	body = strings.ReplaceAll(body, "{RequestIP}", "--")
	body = strings.ReplaceAll(body, "{debugStack}", debugStack)

	if level == "MAIL" {
		// 执行发邮件
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

	} else if level == "SMS" {
		// 执行发短信

	} else if level == "WX" {
		// 执行发微信

	} else if level == "INFO" {
		// 执行记日志
		log.Error("HTTP", errors.New("后端接口异常"), lager.Data{
			"Subject": subject,
			"Body":    body,
		})

		//if f, err := os.OpenFile(config.AppErrorLogName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664); err != nil {
		//	log.Println(err)
		//} else {
		//	errorLogMap := make(map[string]interface{})
		//	errorLogMap["time"] = time.Now().Format("2006/01/02 - 15:04:05")
		//	errorLogMap["info"] = str
		//
		//	errorLogJson, _ := json.Encode(errorLogMap)
		//	_, _ = f.WriteString(errorLogJson + "\n")
		//}
	}
}
