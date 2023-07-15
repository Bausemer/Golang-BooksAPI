package main

import (
	// "errors"
	"github.com/TwiN/go-color"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type book struct{
	ID string 		`json:"id"`
	Title string 	`json:"title"`
	Author string `json:"author"`
	Quantity int 	`json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Harry Potter and the Philosopher's Stone", Author: "J.K.Rowling", Quantity: 2},
	{ID: "2", Title: "Harry Potter and the Chamber of Secrets", Author: "J.K.Rowling", Quantity: 3},
	{ID: "3", Title: "Harry Potter and the Prisoner of Azkaban", Author: "J.K.Rowling", Quantity: 3},
	{ID: "4", Title: "Harry Potter and the Goblet of Fire", Author: "J.K.Rowling", Quantity: 2},
	{ID: "5", Title: "Harry Potter and the Order of the Phoenix", Author: "J.K.Rowling", Quantity: 0},
}

func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, books)
}

func main() {
	port := 8080
	fmt.Println(color.Ize(color.Green,"Api is starting on port: "+strconv.Itoa(port)))
	router := gin.Default()
	router.GET("/books", getBooks)
	router.Run("localhost:"+strconv.Itoa(port))
}