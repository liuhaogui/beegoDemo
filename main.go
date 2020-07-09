package main

import (
	"github.com/astaxie/beego"
	_ "github.com/liuhaogui/beegoDemo/docs"
	_ "github.com/liuhaogui/beegoDemo/routers"
)


func main() {

	// Swagger API
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		//orm.Debug = true
		//loger.InitLog()
	} else {
		//loger.InitLog()
	}

	//beego.RunWithMiddleWares(":8080", metric.PrometheusMiddleWare)
	beego.Run()
}
