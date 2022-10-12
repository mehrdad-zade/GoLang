package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mehrdad-zade/GoLang/3.BookStoreAPI/pkg/models"
	"github.com/mehrdad-zade/GoLang/3.BookStoreAPI/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks) //convert to json
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["boookId"]
	if Id, err := strconv.ParseInt(bookId, 0, 0); err == nil {
		bookDetail, _ := models.GetBookById(Id)
		res, _ := json.Marshal(bookDetail)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
	fmt.Println("Error while parsing the Book ID")
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["boookId"]
	if Id, err := strconv.ParseInt(bookId, 0, 0); err == nil {
		bookDetail := models.DeleteBook(Id)
		res, _ := json.Marshal(bookDetail)
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
	fmt.Println("Error while parsing the Book ID")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	bookDetail, db := models.GetBookById(Id)
	if updateBook.Name != "" {
		bookDetail.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetail.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetail.Publication = updateBook.Publication
	}

	db.Save(&bookDetail)
	res, _ := json.Marshal(bookDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
