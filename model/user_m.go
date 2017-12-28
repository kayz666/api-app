package model

import (
	"time"
	"github.com/kayz666/api-app/utils"
	"regexp"
	"fmt"
	"errors"
	"github.com/satori/go.uuid"
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
type User_m struct{
	Id					int64
	Account    	string    `xorm:"varchar(32) unique" form:"account" json:"account"`
	Email       string	 `xorm:"varchar(64) " form:"email"   json:"email"`
	Telephone   string	 `xorm:"varchar(14) " form:"telephone" json:"telephone"`
	Password   	string    `xorm:"varchar(128)" form:"password" json:"password"`


	Baseinfo			// 基础信息  个人资料相关
	Advancedinfo		//高级信息   系统设置相关
	Deviceinfo   		//设备信息   存储用户设备相关信息
	Authmanageinfo		//权限信息

	Register   time.Time `xorm:"created"`
	Lastlogin  time.Time `xorm:"updated"`

}
//用户信息 基础信息
type Baseinfo struct{
	Name        string  `xorm:"varchar(128)" form:"name" json:"name"`
	Gender      string 	`xorm:"varchar(128)" form:"gender" json:"gender"`
	Age         int
	Iconaddress string   //图标地址

}
//用户信息 系统相关
type Advancedinfo struct{
	Uuid				string 			`xorm:"varchar(64) unique"`
	Apikey				string			`xorm:"varchar(32) unique"`

}


//用户信息  设备相关
type Deviceinfo struct{
	Datashream_table 	string   //用户数据流表名
}
//用户权限管理
type Authmanageinfo struct{

}


//通过邮件方式注册
func Registerwithmail (email string,passwd string)(){

}
//通过手机号码注册
func Registerwithtelphone (num string,passwd string) (){

}


// 检查用户名是否存在
// 参数
//返回 bool  True 存在    Flase 不存在或为空

func AccountIsExist(a string) bool{
	if a==""{
		return false
	}
	re,err := Orm_sys.Get(&User{Account:a})
	if err!=nil{
		utils.LogErr(err.Error())
		return false
	}
	return re
}




// 检查邮箱是否存在
// 参数
//返回 bool  True 存在    Flase 不存在或为空

func EmailIsExist(a string) bool{
	if a==""{
		return false
	}
	re,err := Orm_sys.Get(&User{Email:a})
	if err!=nil{
		utils.LogErr(err.Error())
		return false
	}
	return re
}


// 检查手机号码是否存在
// 参数
//返回 bool  True 存在    Flase 不存在或为空

func TelphoneIsExist(a string) bool{
	if a==""{
		return false
	}
	re,err := Orm_sys.Get(&User{Telephone:a})
	if err!=nil{
		utils.LogErr(err.Error())
		return false
	}
	return re
}

// 检查APIKEY是否存在
// 参数
//返回 bool  True 存在    Flase 不存在或为空
func APIKeyIsExist(a string ) bool{
	if a==""{
		return false
	}
	re,err := Orm_sys.Get(&User{Apikey:a})
	if err!=nil{
		utils.LogErr(err.Error())
		return false
	}
	return re
}

//检查用户名，邮箱，手机号 是否合法
//参数 用户名 至少5位 不超过40位 大小写字母加数字加下划线
//返回 True 合法  Flase 不合法
func CheckUserIsLegal(a string,e string,t string) (bool,bool,bool){
	var ac,em,te bool
	ac=false
	em=false
	te=false
	if a !=""{
		r,_:= regexp.Compile("^[a-zA-Z0-9_]{5,40}$")
		ac=r.MatchString(a)
	}
	if e !=""{
		r,_:= regexp.Compile("^([A-Za-z0-9_.-])+@([A-Za-z0-9_.-])+.([A-Za-z]{2,4})$")
		em=r.MatchString(e)
	}
	if t !=""{
		r,_:= regexp.Compile("^(13[0-9]|14[57]|15[0-35-9]|18[0,5-9]|(17[0-9]))[\\d]{8}$")
		te=r.MatchString(t)
	}
	return ac,em,te
}



// 通过邮件方式注册
func AddNewUserOfEmail(e string,p string) (bool,error) {
	if e == ""{
		return false, errors.New("Email Is Empty")
	}
	if EmailIsExist(e){
		return false, errors.New("Email Is Exist")
	}
	user := &User{}
	user.Account=fmt.Sprintf("User_%s%s",utils.StrToMD5v2(user.Uuid[:15])[:6],uuid.NewV1().String()[:6])
	user.Email=e
	user.Uuid = RandomUUID()
	user.Password = StringToPassword(p)
	user.Apikey = RandomAPIKey()
	user.AuthGroup="ordinary"
	user.Datashream_table = fmt.Sprintf("datastream_%s", uuid.NewV1())
	sql := fmt.Sprintf("CREATE TABLE `%s` %s", user.Datashream_table, datasheet_sql)
	_, err := Orm_data.Query(sql)
	if err != nil {
		user.Datashream_table = fmt.Sprintf("datastream_%s", uuid.NewV1())
		sql := fmt.Sprintf("CREATE TABLE `%s` %s", user.Datashream_table, datasheet_sql)
		_, err = Orm_data.Query(sql)
		if err != nil {
			utils.LogErr(err.Error())
			return false, err
		}
	}
	_, err = Orm_sys.Insert(user)
	if err != nil {
		_, err = Orm_sys.Insert(user)
		if err != nil {
			utils.LogErr(err.Error())
			return false, err
		}
	}
	return true,nil
}

//通过手机号码方式注册
func AddNewUserOfTelphone(t string,p string) (bool,error) {
	if t == ""{
		return false, errors.New("Telphone Is Empty")
	}
	if TelphoneIsExist(t){
		return false, errors.New("Telphone Is Exist")
	}
	user := &User{}
	user.Account=fmt.Sprintf("User_%s%s",utils.StrToMD5v2(user.Uuid[:15])[:6],uuid.NewV1().String()[:6])
	user.Telephone=t
	user.Uuid = RandomUUID()
	user.Password = StringToPassword(p)
	user.Apikey = RandomAPIKey()
	user.AuthGroup="ordinary"
	user.Datashream_table = fmt.Sprintf("datastream_%s", uuid.NewV1())
	sql := fmt.Sprintf("CREATE TABLE `%s` %s", user.Datashream_table, datasheet_sql)
	_, err := Orm_data.Query(sql)
	if err != nil {
		user.Datashream_table = fmt.Sprintf("datastream_%s", uuid.NewV1())
		sql := fmt.Sprintf("CREATE TABLE `%s` %s", user.Datashream_table, datasheet_sql)
		_, err = Orm_data.Query(sql)
		if err != nil {
			utils.LogErr(err.Error())
			return false, err
		}
	}
	_, err = Orm_sys.Insert(user)
	if err != nil {
		_, err = Orm_sys.Insert(user)
		if err != nil {
			utils.LogErr(err.Error())
			return false, err
		}
	}
	return true,nil
}
//验证用户登录消息是否正确
//参数 (account | email | telephone)  password
 //返回 1  user
 //     2  account | email | telephone
 //     3  password
 //		3  err
//func (self *User) UserLoginCheck() (*User,bool,bool,error){
//	user,_,err:= self.GetUser()
//	if err !=nil{
//		return user,false,false,err
//	}
//	if user.Password == utils.StrToMD5v1(self.Password){
//		return user,true,true,nil
//	}
//	return &User{},true,false,errors.New("Password error")
//}




func GetUserOfAccount(s string)(*User,bool){
	user := &User{}
	re,_ :=Orm_sys.Where("account=?",s).Get(user)
	if !re {
		return nil,false
	}
	return user,true
}
func GetUserOfEmail(s string)(*User,bool){
	user := &User{}
	re,_ :=Orm_sys.Where("email=?",s).Get(user)
	if !re {
		return nil,false
	}
	return user,true
}
func GetUserOfTelphone(s string)(*User,bool){
	user := &User{}
	re,_ :=Orm_sys.Where("telphone=?",s).Get(user)
	if !re {
		return nil,false
	}
	return user,true
}
// 根据uuid 获得user结构体
func GetUserOfUUID(s string)(*User,bool){
	user := &User{}
	re,_ :=Orm_sys.Where("uuid=?",s).Get(user)
	if !re {
		return nil,false
	}
	return user,true
}
//根据apikey 获得 user结构体
func GetUserOfAPIKey(s string)(*User,bool){
	user := &User{}
	re,_ :=Orm_sys.Where("apikey=?",s).Get(user)
	if !re {
		return nil,false
	}
	return user,true
}

