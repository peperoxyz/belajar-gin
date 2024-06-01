package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// struct book
type Book struct {
	ID 		string	`json:"id"`
	Title 	string	`json:"title"`
	Stock 	int		`json:"stock"`
	Author 	string	`json:"author"`
}

var Books = []Book{}

func CreateBook(ctx *gin.Context) {
	// variabel book baru untuk nampung data dari formRequest
	var newBook Book

	// karena request berupa raw JSON, maka untuk binding data menggunakan ctx.ShouldBindJSON. namun jika pakai formData, maka pakai ctx.shouldBind
	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// proses append new data
	newBook.ID = fmt.Sprintf("%d", len(Books)+1)
	Books = append(Books, newBook)

	// return response body
	ctx.JSON(http.StatusOK, gin.H {
		"data": newBook,
		"message": "Succeed created book",
	})
}

func GetBooks(ctx *gin.Context) {
	// return response body
	ctx.JSON(http.StatusOK, gin.H {
		"data": Books,
		"message": "Succeed get all books",
	})
}

func GetBook(ctx *gin.Context) {
	// define parameter bookId
	id := ctx.Param("id");
	condition:= false
	var bookData Book

	for i, book := range Books {
		if id == book.ID {
			condition = true
			bookData = Books[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"data": nil,
			"message": fmt.Sprintf("Book with id %v does not exist", id),
		}) 
		return
	}

	// return response body
	ctx.JSON(http.StatusOK, gin.H {
		"data": bookData,
		"message": "Succeed get the book",
	})
}

func DeleteBook(ctx *gin.Context) {
	// define parameter bookId
	id := ctx.Param("id");
	condition:= false

	var bookIndex int

	for i, book := range Books {
		if id == book.ID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"data": nil,
			"message": fmt.Sprintf("Book with id %v does not exist", id),
		}) 
		return
	}

	copy(Books[bookIndex:], Books[bookIndex+1:])
	Books[len(Books)-1] = Book{}
	Books = Books[:len(Books)-1]

	// return response body
	ctx.JSON(http.StatusOK, gin.H {
		"data": nil,
		"message": fmt.Sprintf("Book with id %v has been deleted successfully", id),
	})	
}

func UpdateBook(ctx *gin.Context) {
	// define parameter bookId
	id := ctx.Param("id");
	condition:= false

	var updatedBook Book

	if err := ctx.ShouldBindBodyWithJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range Books {
		if id == book.ID {
			condition = true
			Books[i] = updatedBook
			Books[i].ID = id
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"data": nil,
			"message": fmt.Sprintf("Book with id %v does not exist", id),
		}) 
		return
	}

	// return response body
	convertBookId, _ := strconv.Atoi(id)
	ctx.JSON(http.StatusOK, gin.H {
		"data": Books[convertBookId-1],
		"message": fmt.Sprintf("Book with id %v has been updated successfully", id),
	})

	
}

