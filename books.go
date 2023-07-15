package main

import (
	"errors"
	"strconv"
)

type book struct{
	ID int 		`json:"id"`
	Title string 	`json:"title"`
	Author string `json:"author"`
	Quantity int 	`json:"quantity"`
}

var books = []book{
	{ID: 1, Title: "Harry Potter and the Philosopher's Stone", Author: "J.K.Rowling", Quantity: 2},
	{ID: 2, Title: "Harry Potter and the Chamber of Secrets", Author: "J.K.Rowling", Quantity: 3},
	{ID: 3, Title: "Harry Potter and the Prisoner of Azkaban", Author: "J.K.Rowling", Quantity: 3},
	{ID: 4, Title: "Harry Potter and the Goblet of Fire", Author: "J.K.Rowling", Quantity: 2},
	{ID: 5, Title: "Harry Potter and the Order of the Phoenix", Author: "J.K.Rowling", Quantity: 0},
}

func getBookById(id int) (*book, error){
	
	for i, book := range books {
		if book.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("Book with id " +strconv.Itoa(id) + " not found.")
}