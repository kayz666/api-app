package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	//"github.com/astaxie/beego"
	"github.com/kayz666/api-app/utils"
)

var(
	SysDb *xorm.Engine
)


func DB_init() bool{
	server := "47.97.107.140"
	username := "kay"
	password := "443622fF"
	sys_dbName := "longlee_sys"
	maxIdle := 30
	maxConn := 30
	//data_dbName := "orm_data"
	var err error
	SysDb,err =xorm.NewEngine("mysql",username+":"+password+"@tcp("+server+":3306)/"+sys_dbName+"?charset=utf8")

	if err != nil {
		utils.LogErr(err.Error())
		return false
	}
	//defer SysDb.Close()
	//SysDb.ShowSQL(false)
	SysDb.SetMaxIdleConns(maxIdle)
	SysDb.SetMaxOpenConns(maxConn)
	err = SysDb.Sync(new(User)) //更新模型到数据库
	if err != nil {
		utils.LogErr(err.Error())
		return false
	}
	return true
}

