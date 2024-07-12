package model

type User struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	PhoneNo  string `json:"phone_no"`
	Password string `json:"password"`
}
