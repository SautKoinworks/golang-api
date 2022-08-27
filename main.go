package main

import (
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
	router.POST("/user", postUserHandler)

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
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required,number"`
}

func postBookHandler(c *gin.Context) {

	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "data salah",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}

type UserData struct {
	Name string
	Job  string
	Age  int
}

func postUserHandler(c *gin.Context) {
	var userData UserData

	err := c.ShouldBindJSON(&userData)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"name": userData.Name,
		"job":  userData.Job,
		"age":  userData.Age,
	})
}
