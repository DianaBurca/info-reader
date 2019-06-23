package main

import (
	"github.com/DianaBurca/info-reader/utils"
	"github.com/gin-gonic/gin"
)

func main() {

	utils.EstablishConnection()
	driver := gin.Default()
	driver.GET("/read", utils.ReadHandler)
	driver.Run()
}
