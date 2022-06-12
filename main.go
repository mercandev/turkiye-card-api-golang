package main

import (
	"turkey-card/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/getmockdata", service.GetMockData)

	router.Run("localhost:8542")
}
