package main

import (
	"github.com/astaxie/beego"
	_ "quickStart/routers"
)

func main() {
	//指定静态目录，beego默认设置了static目录，
	//参数：url为访问得路径，path：为当前项目中要设置的文件夹名
	beego.SetStaticPath("/pic", "pic")
	beego.Run()
}
