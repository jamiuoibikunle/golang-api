package routes

import (
	"github.com/gin-gonic/gin"
	"tutorials/controllers"
)

func StateRoute(router *gin.Engine) {
	router.POST("/state", controllers.CreateUser())
	router.GET("/states", controllers.GetUsers())
	router.GET("/states/:stateid", controllers.GetUser())
}