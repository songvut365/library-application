package model

// Book
type NewBook struct {
	Name string `json:"name"`
}

type UpdateBook struct {
	State string `json:"state"`
}

type Book struct {
	ID     *int   `json:"id"`
	BookID int    `json:"book_id"`
	Date   string `json:"date"`
	State  string `json:"state"`
}

// Book Detail
type BookDetail struct {
	ID     *int   `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount" gorm:"default:0"`
}

type UpdateBookDetail struct {
	Name string `json:"name"`
}

// Member
type Member struct {
	ID         *int   `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	NationalID string `json:"national_id"`
}

// Borrow
type Borrow struct {
	ID       *int   `json:"id"`
	BookID   int    `json:"book_id"`
	MemberID int    `json:"member_id"`
	Date     string `json:"date"`
}

type BorrowResult struct {
	ID       *int   `json:"id"`
	BookID   int    `json:"book_id"`
	BookName string `json:"book_name"`
	MemberID int    `json:"member_id"`
	Date     string `json:"date"`
}

type BorrowInput struct {
	BookID     int    `json:"book_id"`
	NationalID string `json:"national_id"`
}
