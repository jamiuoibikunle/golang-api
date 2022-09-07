package routes

import (
	"tutorials/controllers"

	"github.com/gin-gonic/gin"
)

func StateRoute(router *gin.Engine) {
	router.POST("/state", controllers.CreateUser())
	router.GET("/states", controllers.GetUsers())
	router.GET("/states/:stateid", controllers.GetUser())
	router.PUT("/states/:stateid", controllers.EditUser())
	router.DELETE("/states/:stateid", controllers.DeleteUser())
}
