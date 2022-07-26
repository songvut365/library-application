package model

type Book struct {
	ID     *int   `json:"id"`
	BookID int    `json:"book_id"`
	Date   string `json:"date"`
	State  string `json:"state"`
}

type NewBook struct {
	Name string `json:"name"`
}

type UpdateBook struct {
	State string `json:"state"`
}

type BookResponse struct {
	ID    *int   `json:"id"`
	Name  string `json:"name"`
	Date  string `json:"date"`
	State string `json:"state"`
}
