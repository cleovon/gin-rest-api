package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-rest-api/src/app/business"
	"gin-rest-api/src/app/models"
)

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, business.GetAllBooks())
}

func PostBook(c *gin.Context) {
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	business.CreateBook(newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	if a := business.GetBookById(id); a.ID != "" {
		c.IndentedJSON(http.StatusOK, a)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
	}
}

func PutBook(c *gin.Context) {
	id := c.Param("id")

	var updateBook models.Book

	if err := c.BindJSON(&updateBook); err != nil {
		return
	}

	if a := business.UpdateBookById(id, updateBook); a.ID != "" {
		c.IndentedJSON(http.StatusOK, a)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
	}
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	if business.RemoveBook(id) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})

		return
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
	}
}
