package main

import (
	"github.com/gin-gonic/gin"
	"tutorials/configs"
	"tutorials/routes"
)

func main() {
        router := gin.Default()

		configs.ConnectDB()
		routes.StateRoute(router)
  
        router.Run("localhost:6000") 
}