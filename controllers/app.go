package controllers

import (
	"github.com/astaxie/beego/server/web"

	"shorturl-admin/pkg/invoker"
)

type AppController struct {
	web.Controller
}

func (c *AppController) Info() {
	res, _ := invoker.Cfg.String( "name")
	c.Data["json"] = map[string]string{
		"info": res,
	}
	c.ServeJSON()
	return
}
