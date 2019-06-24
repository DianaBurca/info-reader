package main

import (
	"github.com/DianaBurca/info-reader/utils"
	"github.com/gin-gonic/gin"
)

func main() {

	utils.EstablishConnection()
	driver := gin.Default()
	driver.GET("/read", utils.ReadHandler)
	driver.GET("/.well-known/live", utils.Health)
	driver.GET("/.well-known/ready", utils.Health)
	driver.Run()
}
