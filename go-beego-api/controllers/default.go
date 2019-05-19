package controllers

import (
	"github.com/astaxie/beego"
)

// 默认控制器
type DefaultController struct {
	beego.Controller
}

func (m *DefaultController) Get() {
	m.Ctx.WriteString(beego.BConfig.AppName)
}
