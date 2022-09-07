package main

import (
	"tutorials/configs"
	"tutorials/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()
	routes.StateRoute(router)

	router.Use(cors.Default())
	router.RunTLS("localhost:443", "./ssl/cert.pem", "./ssl/key.pem")
}
