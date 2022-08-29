package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"pustaka-api/handler"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=postgres dbname=pustaka-api port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB connection Error")
	}

	fmt.Println("Database connection succeed")
	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/", handler.UserHandler)
	v1.GET("/title", handler.TitleHandler)
	v1.GET("/books/:id/:title", handler.BookHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBookHandler)
	v1.POST("/user", handler.PostUserHandler)

	router.Run(":9999")
}
