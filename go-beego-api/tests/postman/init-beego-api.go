package postman

import (
	"github.com/astaxie/beego"
	"time"
)

type MyController struct {
	beego.Controller
}

func (c *MyController) BeegoApiCurrentTimestamp() {
	c.Ctx.ResponseWriter.Status = 200
	c.Ctx.ResponseWriter.Write([]byte(time.Now().String()))
}
