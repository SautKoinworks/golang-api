package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", userHandler)
	router.GET("/title", titleHandler)
	router.GET("/books/:id/:title", bookHandler)
	router.GET("/query", queryHandler)
	router.POST("/books", postBookHandler)

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

func bookHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

type BookInput struct {
	Title    string
	Price    int
	SubTitle string `json:"sub_title"`
}

func postBookHandler(c *gin.Context) {

	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		log.Fatal(err)
		fmt.Println("error tipe data")
	}

	c.JSON(http.StatusOK, gin.H{
		"title":     bookInput.Title,
		"price":     bookInput.Price,
		"sub_title": bookInput.SubTitle,
	})

}
