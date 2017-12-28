package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kayz666/api-app/routers"
)

func main(){
	router := gin.Default()
	routers.Index(router)
	routers.V1(router)
	routers.V2(router)
	router.Run(":8081")
}