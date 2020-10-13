package routers

import (
	"github.com/astaxie/beego/server/web"
	"shorturl-admin/controllers"
)

func init() {
	web.Router("/", &controllers.AppController{}, "get:Info")
}
