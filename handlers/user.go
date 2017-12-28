package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"strconv"
)

func Getuserlist(c *gin.Context){

}
func Getuserbyid(c *gin.Context){

}

func Adduser(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{
		"status":123,
	})
}



