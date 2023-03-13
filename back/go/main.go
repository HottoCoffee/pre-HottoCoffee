package main

import (
  "github.com/HottoCoffee/HottoCoffee/controller"
  "github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello world"})
	})
	route.Run()

  batchController := controller.BatchController{}
  batchController.GetBatch(1)
}
