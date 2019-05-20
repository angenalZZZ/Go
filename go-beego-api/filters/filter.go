package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
)

func init() {

	// 权限访问控制
	AddPluginOfAuthZ()

	// 自定义过滤
	beego.InsertFilter("/v1/*", beego.BeforeExec, func(ctx *context.Context) {
		basic, auth := "Basic", ctx.Input.Header("Authorization")
		logs.GetLogger(logs.AdapterConsole).Printf(" %s Authorization: %s\n", basic, auth)
	})
}
