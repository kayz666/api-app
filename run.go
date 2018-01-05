package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kayz666/api-app/routers"
	"github.com/kayz666/api-app/model"
)

func main(){
	res:=model.DB_init()
	if !res { return }
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	routers.Index(router)
	routers.V1(router)
	routers.V2(router)

	router.Run(":2345")
}