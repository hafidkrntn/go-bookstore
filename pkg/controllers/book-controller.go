package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hafid/go-bookstore/pkg/models"
	"github.com/hafid/go-bookstore/pkg/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	// get all data di db/models
	newBook := models.GetAllBooks()
	// data to json format
	res, _ := json.Marshal(newBook)
	// from header to content json
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	// get variable in url
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	// change Id to int
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
		// if error while parasing
	}

	// get the data in model/db
	bookDetails, _ := models.GetBookById(ID)

	// change result to json
	res, _ := json.Marshal(bookDetails)

	// setting to header http to content json
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	CreateBook.CreateBook()
	res := map[string]string{
		"message": "Your book was created successfully",
	}
	response, err := json.Marshal(res)
	if err != nil {
		// Tangani kesalahan konversi ke JSON jika diperlukan
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Set header untuk response JSON
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	models.DeleteBook(ID)
	res := map[string]string{
		"message": "Your book was deleted successfully",
	}
	response, err := json.Marshal(res)
	if err != nil {
		// Tangani kesalahan konversi ke JSON jika diperlukan
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// create a new variable for save data book while updating
	var updateBook = &models.Book{}

	// call ParseBody to get data JSON from request body and updating book
	utils.ParseBody(r, updateBook)

	// get id from url
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	// change id string to integer value
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	// get the details books in db / models
	booksDetails, db := models.GetBookById(ID)

	// Check if there is an update on the book name in the data sent.
	// If there is, then update the book data with the new value.
	if updateBook.Name != "" {
		booksDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		booksDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		booksDetails.Publication = updateBook.Publication
	}

	// save data or update data in db / models
	db.Save(&booksDetails)
	res := map[string]string{
		"message": "Your book was updated successfully",
	}
	response, err := json.Marshal(res)
	if err != nil {
		// Tangani kesalahan konversi ke JSON jika diperlukan
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
