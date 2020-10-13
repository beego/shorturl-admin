package main

import (
	_ "github.com/astaxie/beego/core/config/toml"
	"github.com/astaxie/beego/server/web"
	"shorturl-admin/pkg/invoker"
	"shorturl-admin/routers"
)

func main() {
	invoker.Init()
	routers.Init()
	web.Run()
}
