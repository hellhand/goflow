package model

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	BirthDate int64  `json:"birthDate"`
}
