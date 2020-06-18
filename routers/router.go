package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"quickStart/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//匹配http://localhost:8080/hello/1 指定HTTP请求，处理方法方法及RESTful规则
	beego.Router("/hello/?:id", &controllers.HelloController{}, "post:Hello")
	beego.Router("/not/:id", &controllers.HelloController{}, "post:Hello")
	//支持get请求的路由，类似的还有post put等
	beego.Get("/get", func(context *context.Context) {
		context.Output.Body([]byte("基础路由Get()"))
	})
	//响应任何HTTP请求的路由
	beego.Any("/any", func(c *context.Context) {
		c.Output.Body([]byte("响应任何HTTP请求的路由Any()"))
	})
	//正则匹配
	beego.Router("/api/:id([0-9]+)", &controllers.HelloController{}, "get:Hello")
	beego.Router("/api/:id([\\w]+)", &controllers.HelloController{}, "get:Hello")
	//http://localhost:8080/download/abc/msr.html，可以匹配任意的扩展名".html .png"等，也可以指定扩展名
	//download到扩展名之间可以匹配任意的路径
	beego.Router("/download/*.*", &controllers.HelloController{})
	//全匹配 all后面可以匹配任意的路径，包括扩展名
	beego.Router("/all/*", &controllers.HelloController{})
	//指定参数类型 框架实现正则([0-9]+) ([\\w]+)
	beego.Router("/int/:id:int", &controllers.HelloController{})
	beego.Router("/string/:id:string", &controllers.HelloController{})
	//带前缀正则匹配
	beego.Router("/prefix/post_:id([0-9]+).html", &controllers.HelloController{})

}
