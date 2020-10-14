package main

import (
	_ "github.com/astaxie/beego/core/config/toml"
	"github.com/astaxie/beego/core/logs"
	"github.com/astaxie/beego/server/web"

	"shorturl-admin/pkg/invoker"
	"shorturl-admin/routers"
)

func main() {

	err := web.LoadAppConfig("toml", "config/local.toml")
	if err != nil {
		logs.Error("could not load the app config: ", err)
	}

	invoker.Init()
	routers.Init()
	web.Run()

}
