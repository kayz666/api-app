package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"strconv"
	"github.com/kayz666/api-app/model"
	"github.com/kayz666/api-app/utils"

)



//@Titer  register a new user
//
//@router /v1/user/register?mode=<>  [post]
func UserReg(c *gin.Context){
	mode:=c.DefaultQuery("mode","email")
	if mode=="email"{
		email:=c.PostForm("email")
		passwd:=c.PostForm("password")
		user,err :=model.RegUserOfEmail(email,passwd)
		if err !=nil{
			c.JSON(http.StatusOK,gin.H{
				"ecode":10101,
				"success": false,
				"error":err.Error(),
				"return":gin.H{
					"link":"",
					"email":"",
				},
			})
			return
		}
		err =utils.SSendRegisterEmail(email,user.Authentication_data)
		if err != nil{
			c.JSON(http.StatusOK,gin.H{
				"ecode":10102,
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
				"link":user.Authentication_data,
				"email":user.Email,
			},
		})
		return
	}else if mode == "telphone"{
		c.JSON(http.StatusOK,gin.H{
			"ecode":10103,
			"success": false,
			"error":"telphone ",
			"return":gin.H{
				"code":"",
			},
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"ecode":10104,
		"success": false,
		"error":"unkown mode",
		"return":gin.H{
			"link":"",
		},
	})
}

//注册认证
func UserVali(c *gin.Context){
	email:= c.PostForm("email")
	code := c.PostForm("code")
	user :=model.GetUserWithEmail(email)
	if user.Authentication_data==code{
		err :=user.SetUserRegisterSuccess()
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"ecode":10105,
				"success": false,
				"error":err.Error(),
				"return":gin.H{
				},
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"success": true,
			"error":nil,
			"return":gin.H{
				"email":user.Email,
			},
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":10202,
		"success": false,
		"error":"",
		"return":gin.H{
		},
	})
}


//用户登陆
func UserLogin(c *gin.Context){

}

