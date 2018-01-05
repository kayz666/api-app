package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"strconv"
	"github.com/kayz666/api-app/model"
	"github.com/kayz666/api-app/utils"

)
//@Titer  Get all user infomation
//@Param
//@router / [get]
func Getuserall(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"status":"11/100",
	})
}
//@Titer  Get user by uid
//
//@router /:id  [get]
func Getuser(c *gin.Context){

}

//@Titer  register a new user
//
//@router /  [post]
func Reguser(c *gin.Context) {
	mode:=c.DefaultQuery("mode","email")
	email:=c.PostForm("email")
	//telphone:=c.Query("telphone")
	passwd:=c.PostForm("password")
	if mode=="email"{
		user,err :=model.RegUserOfEmail(email,passwd)
		if err !=nil{
			c.JSON(http.StatusOK,gin.H{
				"code":10101,
				"success": false,
				"error":err.Error(),
				"return":gin.H{
					"link":"",
					"email":"",
				},
			})
			return
		}
		err =utils.SSendRegisterEmail(email,user.Setting.Authentication_data)
		if err != nil{
			c.JSON(http.StatusOK,gin.H{
				"code":10102,
				"success": false,
				"error":err.Error(),
				"return":gin.H{
					"link":"",
					"email":"",
				},
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"success": true,
			"error":nil,
			"return":gin.H{
				"link":user.Setting.Authentication_data,
				"email":user.Email,
			},
		})
		return
	}else if mode == "telphone"{
		c.JSON(http.StatusOK,gin.H{
			"code":10103,
			"success": false,
			"error":"telphone ",
			"return":gin.H{
				"code":"",
			},
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":10104,
		"success": false,
		"error":"unkown mode",
		"return":gin.H{
			"link":"",
		},
	})
}



