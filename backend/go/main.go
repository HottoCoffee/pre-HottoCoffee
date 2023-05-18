package main

import (
	"fmt"
	"github.com/HottoCoffee/HottoCoffee/controller"
	"github.com/HottoCoffee/HottoCoffee/infrastructure"
	"github.com/HottoCoffee/HottoCoffee/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func main() {
	route := SetUp()
	if route == nil || route.Run() != nil {
		panic("Failed to boot app")
	}
}

func SetUp() *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load environment file")
		return nil
	}
	dialector := mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/hottocoffee?parseTime=True", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST")))
	db, dbErr := gorm.Open(dialector, &gorm.Config{})
	if dbErr != nil {
		fmt.Println("Failed to connect DB")
		return nil
	}
	br := infrastructure.NewBatchRepository(db)
	hr := infrastructure.NewHistoryRepositoryImpl(*db, br)

	route := gin.Default()
	route.GET("/api/batch", func(c *gin.Context) {
		query := c.Query("query")
		bp := controller.NewBatchPresenter(c)
		usecase.NewGetBatchListUsecase(br, &bp).Execute(query)
	})
	route.GET("/api/batch/:id", func(c *gin.Context) {
		bp := controller.NewBatchPresenter(c)
		usecase.NewGetBatchUsecase(br, &bp).Execute(c.Param("id"))
	})

	route.GET("/api/batch/:id/history", func(c *gin.Context) {
		hp := controller.NewHistoryPresenter(c)
		usecase.NewGetHistoryListUsecase(hr, hp).Execute(c.Param("id"))
	})
	route.GET("/api/batch/:id/history/:historyId", func(c *gin.Context) {
		hp := controller.NewHistoryPresenter(c)
		usecase.NewGetHistoryUsecase(hr, hp).Execute(c.Param("id"), c.Param("historyId"))
	})

	route.GET("/api/calendar", func(c *gin.Context) {
		startDatetimeInput := c.Query("start_date")
		endDatetimeInput := c.Query("end_date")
		hp := controller.NewHistoryPresenter(c)
		usecase.NewGetCalendarUsecase(hp, hr).Execute(startDatetimeInput, endDatetimeInput)
	})

	return route
}
