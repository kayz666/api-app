package model

import (
	"time"
	"errors"
	"github.com/kayz666/api-app/utils"
	//"fmt"
	"fmt"
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


	//Profile					`xorm:"extends"`// 基础信息  个人资料相关
		Name        string  `xorm:"varchar(128)" form:"name" json:"name"`
		Gender      string 	`xorm:"varchar(128)" form:"gender" json:"gender"`
		Age         int
		Iconaddress string   //图标地址
	//Setting					`xorm:"extends"`//高级信息   系统设置相关
		Uuid				string 			`xorm:"varchar(64)"`
		Apikey				string			`xorm:"varchar(64)"`
		Isauthentication	int				`xorm:"default(0)"`
		Authentication_mode	string
		Authentication_data	string 			`xorm:"varchar(1024)"`
	//Deviceinfo	   			`xorm:"extends"`//设备信息   存储用户设备相关信息
		Datashream_table 	string   //用户数据流表名
	//Authmanageinfo	 		`xorm:"extends"`//权限信息
		AuthGroup		[]string
	//登陆信息
		Login_tokenkey		string			`xorm:"varchar(64)"`
		Login_Outtime		time.Time



	Register   time.Time `xorm:"created"`
	Lastlogin  time.Time `xorm:"updated"`

}



//用户认证成功  初始化用户结构体数据
func (u *User) SetUserRegisterSuccess() error{
	user :=&User{}
	user.Uuid=RandomUUID()
	user.Apikey=RandomAPIKey()
	user.Datashream_table=fmt.Sprintf("ds_%s",user.Uuid)
	sql := fmt.Sprintf("CREATE TABLE `%s` %s", user.Datashream_table, datasheet_sql)
	_, err := DataDb.Query(sql)
	if err!= nil{
		utils.LogErr(err.Error())
		return err
	}
	user.Isauthentication =1
	_,err = SysDb.Id(u.Id).Update(user)
	//_,err := SysDb.Table(new(User)).Id(u.Id).Update(map[string]interface{}{"isauthentication":1,
	//
	//																		})
	if err!= nil {
		utils.LogErr(err.Error())
		return err
	}
	return err
}
// 得到user 结构体
func (u *User) GetUser (){
	_,err :=SysDb.Get(u)
	if err != nil {
		utils.LogDebug(err.Error())
		u= nil
	}
}









// 通过 Email方式 初始化一个注册数据到数据库
func RegUserOfEmail(e string,p string)(*User,error){

	if e==""{return nil,errors.New("Email is empty!")}
	if IsEmailExist(e){return nil,errors.New("Email is exist")}
	user:= &User{Email:e,Isauthentication:0,Authentication_mode:"EMAIL",Authentication_data:RandomSN()}
	user.Password=MakePasswd(p)
	user.Account=RandomAccount()
	//fmt.Println(user)
	_,err:=SysDb.Insert(user)
	if err!=nil{
		utils.LogDebug(err.Error())
		return nil,err
	}

	return user,nil
}


func IsEmailExist(e string) bool{
	re,_:=SysDb.Get(&User{Email:e})
	return re
}
func IsTelphoneExist(t string) bool{
	re,_:=SysDb.Get(&User{Telephone:t})
	return re
}

func GetUserWithEmail(e string) *User{
	user := &User{Email:e}
	user.GetUser()
	return user
}

