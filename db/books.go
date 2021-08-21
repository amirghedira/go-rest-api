package db

import (
	"github.com/amirghedira/go-rest-api/models"
)

var Books = append([]models.Book{}, models.Book{Id: "1", Title: "book1", Author: &models.Author{Firstname: "amir", Lastname: "ghedira"}}, models.Book{Id: "3", Title: "book3", Author: &models.Author{Firstname: "steeve", Lastname: "smith"}}, models.Book{Id: "2", Title: "book2", Author: &models.Author{Firstname: "ahmed", Lastname: "kerkni"}})
