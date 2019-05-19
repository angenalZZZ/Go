package controllers

import (
	"github.com/astaxie/beego"
)

type DefaultController struct {
	beego.Controller
}

func (m *DefaultController) Get() {
	m.Ctx.WriteString(beego.BConfig.AppName)
}
