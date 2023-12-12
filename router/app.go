package router

import (
	"OnlineVideo/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.POST("/user/createUser", service.CreateUser)
	return r
}
