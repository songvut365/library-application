package model

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

type BorrowResponse struct {
	ID       *int   `json:"id"`
	BookID   int    `json:"book_id"`
	BookName string `json:"book_name"`
	MemberID int    `json:"member_id"`
	Date     string `json:"date"`
}
