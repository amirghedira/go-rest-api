package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/amirghedira/go-rest-api/db"
	"github.com/amirghedira/go-rest-api/models"
	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.Books)

}

func GetBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range db.Books {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(404)
	message := make(map[string]string)
	message["error"] = "Book not found"
	json.NewEncoder(w).Encode(message)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Intn(10000000000))
	db.Books = append(db.Books, book)
	json.NewEncoder(w).Encode(book)

}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	params := mux.Vars(r)

	json.NewDecoder(r.Body).Decode(&book)
	for i, item := range db.Books {
		if item.Id == params["id"] {
			book.Id = db.Books[i].Id
			db.Books[i] = book
			json.NewEncoder(w).Encode(db.Books[i])
			return

		}
	}
	w.WriteHeader(404)
	message := make(map[string]string)
	message["error"] = "Book not found"
	json.NewEncoder(w).Encode(message)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range db.Books {
		if item.Id == params["id"] {

			db.Books = append(db.Books[:i], db.Books[i+1:]...)
			responseMessage := make(map[string]string)
			responseMessage["message"] = "successfully delete book with id " + params["id"]
			json.NewEncoder(w).Encode(responseMessage)
			return

		}
	}
	w.WriteHeader(404)
	message := make(map[string]string)
	message["error"] = "Book not found"
	json.NewEncoder(w).Encode(message)

}
