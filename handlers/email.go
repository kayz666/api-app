package handlers

import (
	"github.com/gin-gonic/gin"
	//"strconv"
	"github.com/kayz666/api-app/model"
	"net/http"
	"fmt"
	//"github.com/kayz666/api-app/utils"
)

//@router /v1/email/:email/   [POST]
func PostEmail(c *gin.Context){
	email:= c.Param("email")
	code := c.PostForm("code")
	user :=model.FindUserWithEmail(email)
	fmt.Println(user)
	if user.Setting.Authentication_data==code{
		user.SetUserRegisterSuccess()
		fmt.Println(user)
		c.JSON(http.StatusOK,gin.H{
			"code":10201,
			"success": true,
			"error":nil,
			"return":gin.H{
			},
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code":10202,
		"success": false,
		"error":"",
		"return":gin.H{
			"link":"",
			"email":"",
		},
	})
}


func GetEmailAll(c *gin.Context){

}

func GetEmail(c *gin.Context){

}