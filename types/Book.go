package types

type Book struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}
