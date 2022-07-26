package model

type Member struct {
	ID         *int   `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	NationalID string `json:"national_id"`
}
