package main

import (
	"fmt"
	"github.com/HottoCoffee/HottoCoffee/controller"
	"github.com/HottoCoffee/HottoCoffee/infrastructure"
	"github.com/HottoCoffee/HottoCoffee/usecase"
	"github.com/gin-contrib/cors"
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

	route := gin.Default()
	route.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("ALLOW_ORIGIN")},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
	}))
	route.GET("/api/batch", func(c *gin.Context) {
		query := c.Query("query")
		bp := controller.NewBatchPresenter(c)
		usecase.NewGetBatchListUsecase(br, &bp).Execute(query)
	})
	route.GET("/api/batch/:id", func(c *gin.Context) {
		bp := controller.NewBatchPresenter(c)
		usecase.NewGetBatchUsecase(br, &bp).Execute(c.Param("id"))
	})
	route.POST("/api/batch", func(c *gin.Context) {
		bp := controller.NewBatchPresenter(c)
		input := usecase.BatchInput{}
		if c.ShouldBind(input) != nil {
			bp.SendInvalidRequestResponse("Invalid format")
		}
		usecase.NewCreateBatchUsecase(br, &bp).Execute(input)
	})
	return route
}
