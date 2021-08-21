package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	_ "github.com/amirghedira/go-rest-api/types"
	"github.com/gorilla/mux"
)

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
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

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Intn(10000000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	params := mux.Vars(r)

	json.NewDecoder(r.Body).Decode(&book)
	for i, item := range books {
		if item.Id == params["id"] {
			book.Id = books[i].Id
			books[i] = book
			json.NewEncoder(w).Encode(books[i])
			return

		}
	}
	w.WriteHeader(404)
	message := make(map[string]string)
	message["error"] = "Book not found"
	json.NewEncoder(w).Encode(message)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, item := range books {
		println(i)
		if item.Id == params["id"] {

			books = append(books[:i], books[i+1:]...)
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

func main() {

	books = append(books, Book{Id: "1", Title: "book1", Author: &Author{Firstname: "amir", Lastname: "ghedira"}})
	books = append(books, Book{Id: "2", Title: "book2", Author: &Author{Firstname: "ahmed", Lastname: "kerkni"}})
	books = append(books, Book{Id: "3", Title: "book3", Author: &Author{Firstname: "steeve", Lastname: "smith"}})
	r := mux.NewRouter()
	r.HandleFunc("/book", getBooks).Methods("GET")
	r.HandleFunc("/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/book/{id}", updateBook).Methods("PATCH")
	r.HandleFunc("/book/{id}", deleteBook).Methods("DELETE")
	r.HandleFunc("/book", createBook).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", r))
}
