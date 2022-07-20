package model

type NewBook struct {
	Name string `json:"name"`
}

type UpdateBook struct {
	State string `json:"state"`
}

type BookDetail struct {
	ID     *int   `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount" gorm:"default:0"`
}

type Book struct {
	ID     *int   `json:"id"`
	BookID int    `json:"book_id"`
	Date   string `json:"date"`
	State  string `json:"state"`
}

type Member struct {
	ID         *int   `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	NationalID string `json:"national_id"`
}

type Borrow struct {
	ID       *int   `json:"id"`
	BookID   int    `json:"book_id"`
	MemberID int    `json:"member_id"`
	Date     string `json:"date"`
}

type BorrowInput struct {
	BookID     int    `json:"book_id"`
	NationalID string `json:"national_id"`
}
