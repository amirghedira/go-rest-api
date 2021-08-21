package models

type Book struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}
