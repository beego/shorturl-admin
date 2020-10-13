package invoker

import (
	"github.com/astaxie/beego/core/config"
	"github.com/astaxie/beego/core/logs"
)

var (
	Cfg config.Configer
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
}
