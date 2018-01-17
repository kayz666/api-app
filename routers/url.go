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

		v1.POST		("/user/register",			handlers.UserReg)   //用户注册
		v1.POST		("/user/validate",			handlers.UserVali)	//用户验证
		v1.POST		("/user/login",			handlers.UserLogin)

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

