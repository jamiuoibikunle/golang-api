package main

import (
	"tutorials/configs"
	"tutorials/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()
	routes.StateRoute(router)

	router.RunTLS("localhost:6000", "./ssl/cert.pem", "./ssl/key.unencrypted.pem")
}
