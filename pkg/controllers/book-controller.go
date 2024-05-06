package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Shreyank031/go-bookstore/pkg/models"
	"github.com/Shreyank031/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewBook models.Book

// get books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks() //store getAllBooks from database to newBooks
	res, _ := json.Marshal(newBooks) //res = json version of newBooks which we found in database
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res) //Send something to the frontend/postman

}
func GetBookById(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r) //access our request to get the id from the user
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0) //convert from string to int
	if err != nil {
		fmt.Println("Error while parsing!")
	}
	bookDetails, _ := models.GetBookById(ID) //all the above code to send ID in database to search
	//i don't want to use the db variable right now so iam using the blank character-> see models
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	CreateBookk := &models.Book{}
	utils.ParseBody(r, CreateBookk) //r-> the book which is sent by the user as response and sore it inside the interface CreateBook. refer utils.ParseBody
	b := CreateBookk.CreateBook()   //We sent the parsed/unMarshalled json to databse
	res, _ := json.Marshal(b)       //we are converting the the record to marshalled json. response
	//	for user
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	Delete_Book := models.DeleteBookById(ID)
	res, _ := json.Marshal(Delete_Book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	//step 1 -> To parse the json to unmarshalled json  comming from the user as a request
	utils.ParseBody(r, updateBook)
	//step 2 -> id from request and convert it to int
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
