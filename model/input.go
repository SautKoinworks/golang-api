package model

import "encoding/json"

type BookInput struct {
	Title string      `json:"title" binding:"required"`
	Price json.Number `json:"price" binding:"required,number"`
}

type UserData struct {
	Name string
	Job  string
	Age  int
}
