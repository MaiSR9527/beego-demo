package routers

import (
	"github.com/astaxie/beego"
	"quickStart/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
