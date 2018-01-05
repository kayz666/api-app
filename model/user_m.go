package model

import (
	"time"
	"errors"
	"github.com/kayz666/api-app/utils"
	//"fmt"
)

// 数据流 结构体
//type Datastream struct {
//	Id				int64
//	Deviceid		int64
//	Name 			string
//	Type 			string
//	Key             string
//	Value			string
//	Time			time.Time     `xorm:"created"`
//}
//创建 数据流 表 SQL结构命令
const datasheet_sql  = "(`id`  bigint NOT NULL AUTO_INCREMENT ," +
						"`deviceid`  int NULL ," +
						"`name`  varchar(64) NULL,"+
						"`type`  varchar(255) NULL ," +
						"`key`  varchar(255) NULL," +
						"`value`  varchar(255) NULL," +
						"`time`  datetime NULL ,PRIMARY KEY (`id`)) DEFAULT CHARSET utf8"


//用户信息结构体
type User struct{
	Id			int64
	Account    	string    	`xorm:"varchar(32) unique" form:"account" json:"account"`
	Email       string	 	`xorm:"varchar(64) " form:"email"   json:"email"`
	Telephone   string	 	`xorm:"varchar(14) " form:"telephone" json:"telephone"`
	Password   	string    	`xorm:"varchar(128)" form:"password" json:"password"`


	Profile					`xorm:"extends"`// 基础信息  个人资料相关
	Setting					`xorm:"extends"`//高级信息   系统设置相关
	Deviceinfo	   			`xorm:"extends"`//设备信息   存储用户设备相关信息
	Authmanageinfo	 		`xorm:"extends"`//权限信息


	Register   time.Time `xorm:"created"`
	Lastlogin  time.Time `xorm:"updated"`

}
//用户信息 基础信息
type Profile struct{
	Name        string  `xorm:"varchar(128)" form:"name" json:"name"`
	Gender      string 	`xorm:"varchar(128)" form:"gender" json:"gender"`
	Age         int
	Iconaddress string   //图标地址

}
//用户信息 系统相关
type Setting struct{
	//UserId 				int64  `xorm:"index"`
	Uuid				string 			`xorm:"varchar(64)"`
	Apikey				string			`xorm:"varchar(32)"`
	Isauthentication	bool			`xorm:"default(0)"`
	Authentication_mode	string
	Authentication_data	string 			`xorm:"varchar(1024)"`
}

//用户信息  设备相关
type Deviceinfo struct{
	Datashream_table 	string   //用户数据流表名
}
//用户权限管理
type Authmanageinfo struct{
	AuthGroup		[]string
}


type user_handle interface{
	SetEmailRegisterFlag(b bool)
}

func (u *User) SetUserRegisterSuccess() error{
	user :=&User{}
	user.Setting.Isauthentication=true
	_,err := SysDb.Table(new(User)).Id(u.Id).Update(map[string]interface{}{"isauthentication":1})
	if err!= nil {
		utils.LogDebug(err.Error())
	}
	return err
}

func (u *User) GetUser (){
	_,err :=SysDb.Get(u)
	if err != nil {
		utils.LogDebug(err.Error())
		u= nil
	}
}










func RegUserOfEmail(e string,p string)(*User,error){

	if e==""{return nil,errors.New("Email is empty!")}
	if IsEmailExist(e){return nil,errors.New("Email is exist")}
	user:= &User{Email:e}
	user.Password=MakePasswd(p)
	user.Account=RandomAccount()
	user.Setting =Setting{Isauthentication:false,Authentication_mode:"EMAIL",Authentication_data:RandomSN()}
	//fmt.Println(user)
	_,err:=SysDb.Insert(user)
	if err!=nil{
		utils.LogDebug(err.Error())
		return nil,err
	}

	return user,nil
}

func AddUser(){

}


func IsEmailExist(e string) bool{
	re,_:=SysDb.Get(&User{Email:e})
	return re
}
func IsTelphoneExist(t string) bool{
	re,_:=SysDb.Get(&User{Telephone:t})
	return re
}

func FindUserWithEmail(e string) *User{
	user := &User{Email:e}
	user.GetUser()
	return user
}

