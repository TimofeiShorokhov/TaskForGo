package main

import (
	"github.com/TimofeiShorokhov/TaskForGo/Controller"
	_ "github.com/TimofeiShorokhov/TaskForGo/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Test Task
// @description Realized GET and POST methods
func main() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.SetTrustedProxies([]string{"192.168.1.2"})
	router.GET("/requests", Controller.GetRequests)
	router.GET("/responses", Controller.GetResponses)
	router.POST("/user", Controller.PostUser)
	router.Run("localhost:8080")

}
