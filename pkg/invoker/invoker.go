package invoker

import (
	"github.com/astaxie/beego/client/orm"
	"github.com/astaxie/beego/core/config"
	"github.com/astaxie/beego/core/logs"
	"github.com/beego/invoker"
	"github.com/beego/invoker/orm/mysql"
)

var (
	Cfg config.Configer
	Db  orm.Ormer
)

var (
	ConfigFile = "./config/local.toml"
	err        error
)

func Init() {
	Cfg, err = config.NewConfig("toml", ConfigFile)
	if err != nil {
		logs.Critical("An error occurred:", err)
		panic(err)
	}
	invoker.Init(Cfg)
	Db, err = mysql.Invoker("beego.mysql").Build()
	if err != nil {
		logs.Critical("An error occurred:", err)
		panic(err)
	}
}
