package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/kayz666/api-app/handlers"
	"github.com/kayz666/api-app/middlewares"
	"net/http"
)
func Index(router *gin.Engine){
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"index")
	})
}

func V1(router *gin.Engine) {
	router.GET("/v1", func(c *gin.Context) {
		c.String(http.StatusOK,"v1")
	})
	v1 := router.Group("/v1",middlewares.MiddleWare_Auth())
	{
		v1.GET("/user",handlers.Getuserlist)   //获得用户列表
		v1.GET("/user/:id/*var",handlers.Getuserbyid)  //获得某用户信息
		v1.POST("/user",handlers.Adduser)  //创建新用户
		v1.PUT("/user/:id")   //更新用户
		v1.PATCH("/user/:id/*var")  //更新用户，指定部分
		v1.DELETE("/user/:id")  //删除用户

		v1.POST("/login/:mode",)  //用户获取登陆信息
	}

}
func V2(router *gin.Engine) {
	router.GET("/v2",func(c *gin.Context){
		c.String(http.StatusOK,"v2")
	})
	v2 := router.Group("/v2")
	{
		v2.GET("/user",func(c *gin.Context){
			c.String(http.StatusOK,"v2/user")
		})
		user:=v2.Group("/user")
		{
			user.GET("/*id",func(c *gin.Context){
				id := c.Param("id")
				c.String(http.StatusOK,id)
			})
		}

	}

}

