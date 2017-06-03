package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "lethergo.com"
	c.Data["Email"] = "yggamail@gmail.com"
	c.TplName = "index.html"
}
