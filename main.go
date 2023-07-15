package main

import (
	"github.com/TwiN/go-color"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)


func getBooks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, books)
}

func createBook(context *gin.Context){
	var newBook book
	
	if err := context.BindJSON(&newBook); err != nil {
		return
	}

	highestId := -1
	for _, book := range books {
		if book.ID > highestId {
			highestId = book.ID + 1
		}
	}

	newBook.ID = highestId

	books = append(books, newBook)
	context.IndentedJSON(http.StatusCreated, newBook)
}

func getBook(context *gin.Context){
	id := context.Param("id")
	idAsInt, _ := strconv.Atoi(id)
	book, err := getBookById(idAsInt)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return 
	}

	context.IndentedJSON(http.StatusOK, book)
}

func main() {
	port := 8080
	fmt.Println(color.Ize(color.Green,"Api is starting on port: "+strconv.Itoa(port)))
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBook)
	router.POST("/books", createBook)
	router.Run("localhost:"+strconv.Itoa(port))
}