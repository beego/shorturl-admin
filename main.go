package main

import (
	_ "github.com/astaxie/beego/core/config/toml"
	"github.com/astaxie/beego/server/web"
	"shorturl-admin/pkg/invoker"
)

func main() {
	invoker.Init()
	web.Run()
}
