package business

import (
	"gin-rest-api/src/app/models"
)

// albums slice to seed record album data.
var books = []models.Book{
	{ID: "1", Title: "Revolução dos Bichos", Author: "George Orwell", Price: 56.99},
	{ID: "2", Title: "Harry Potter", Author: "JK Rowling", Price: 17.99},
	{ID: "3", Title: "O Homem da Máscara de Ferro", Author: "Alexandre Dumas", Price: 39.99},
}

func CreateBook(newBook models.Book) {
	books = append(books, newBook)
}

func GetAllBooks() []models.Book {
	return books
}

func GetBookById(id string) models.Book {
	for _, a := range books {
		if a.ID == id {
			return a
		}
	}
	return models.Book{}
}

func UpdateBookById(id string, updatedAlbum models.Book) models.Book {
	for i, a := range books {
		if a.ID == id {
			books[i] = updatedAlbum
			return books[i]
		}
	}
	return models.Book{}
}

func RemoveBook(id string) bool {
	for i, a := range books {
		if a.ID == id {

			books = append(books[:i], books[i+1:]...)

			return true
		}
	}
	return false
}
