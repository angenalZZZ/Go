package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

func init() {
	beego.InsertFilter("/v1/*", beego.BeforeExec, func(ctx *context.Context) {
		basic, auth := "Basic", ctx.Input.Header("Authorization")
		if auth != "" {
			logs.GetLogger(logs.AdapterConsole).Printf(" %s Authorization: %s\n", basic, auth)
		}
	})
}
