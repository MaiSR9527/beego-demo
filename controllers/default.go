package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {

	beego.Controller
}

type HelloController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
func (h *HelloController) Get() {
	h.Data["Name"] = "hello maishuren!"
	h.TplName = "hello.html"
}
func (h *HelloController) Hello() {
	id := h.Ctx.Input.Param(":id")
	h.Data["Name"] = "hello maishuren!Welcome to use beego!id=" + id
	h.TplName = "hello.html"
}
