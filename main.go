package main

import (
	"github.com/gin-gonic/gin"

	"pustaka-api/handler"
)

func main() {
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
