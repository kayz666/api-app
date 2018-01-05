package utils

import (
	//"../utils"
	//"../beego/config"
	"github.com/astaxie/beego/config"
	"fmt"
)

var Config config.Configer

func init( ){
	Conf, err := config.NewConfig("ini", "./conf/app.conf")
	if err != nil {
		fmt.Println(err.Error())
	}
	Config = Conf
}
