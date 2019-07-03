package model

// Author @felixgravila

// Book is the main object
type Book struct {
	ID     int    `json:"id"`
	ISBN   string `json:"isbn"`
	Title  string `json:"title"`
	Author int    `json:"author"`
}

// Author is the person who writes the books
type Author struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
