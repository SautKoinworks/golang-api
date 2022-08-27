package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"pustaka-api/model"
)

func UserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "saut sihotang",
		"job":  "backend enginner",
		"age":  23,
	})
}

func TitleHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "belajar golang api",
	})
}

func BookHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

func PostBookHandler(c *gin.Context) {

	var bookInput model.BookInput

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

func PostUserHandler(c *gin.Context) {
	var userData model.UserData

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
