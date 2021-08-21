package routes

import (
	"github.com/amirghedira/go-rest-api/controllers"
	"github.com/gorilla/mux"
)

func BookApi(BookRoutes *mux.Router) {
	BookRoutes.HandleFunc("/", controllers.GetBooks).Methods("GET")
	BookRoutes.HandleFunc("/{id}", controllers.GetBook).Methods("GET")
	BookRoutes.HandleFunc("/{id}", controllers.UpdateBook).Methods("PATCH")
	BookRoutes.HandleFunc("/{id}", controllers.DeleteBook).Methods("DELETE")
	BookRoutes.HandleFunc("/", controllers.CreateBook).Methods("POST")
}
