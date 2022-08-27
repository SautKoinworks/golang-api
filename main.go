package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/", userHandler)
	v1.GET("/title", titleHandler)
	v1.GET("/books/:id/:title", bookHandler)
	v1.GET("/query", queryHandler)
	v1.POST("/books", postBookHandler)
	v1.POST("/user", postUserHandler)

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
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
}

func postBookHandler(c *gin.Context) {

	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		errorMassages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMassage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMassages = append(errorMassages, errorMassage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMassages,
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
