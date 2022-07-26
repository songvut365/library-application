package model

// Book Detail
type BookDetail struct {
	ID     *int   `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount" gorm:"default:0"`
}

type UpdateBookDetail struct {
	Name string `json:"name"`
}
