package main

import (
	"fmt"
	"github.com/HottoCoffee/HottoCoffee/controller"
	"github.com/HottoCoffee/HottoCoffee/infrastructure"
	"github.com/HottoCoffee/HottoCoffee/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// DI
	dialector := mysql.Open("root:root@tcp(127.0.0.1)/hottocoffee?parseTime=True")
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect DB")
		return
	}
	br := infrastructure.NewBatchRepository(db)
	bu := usecase.NewGetBatchUsecase(br)

	route := gin.Default()
	route.GET("/api/batch/:id", func(c *gin.Context) {
		bp := controller.NewBatchPresenter(c)
		bc := controller.NewBatchController(bu, &bp)
		bc.GetBatch(c.Param("id"))
	})
	route.Run()
}
