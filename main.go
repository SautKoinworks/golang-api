package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", userHandler)
	router.GET("/title", titleHandler)

	router.Run(":9999")
}

func userHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "saut sihotang",
		"job":  "backend enginner",
		"age":  23,
	})
}

func titleHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "belajar golang api",
	})
}