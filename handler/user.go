package handler

import (
	"fmt"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// 	http.Error(w, "Only post methord allowed", http.StatusBadRequest)
	// 	return
	// }
	if r.URL.Path != "/register" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	FullName := r.FormValue("full_name")
	email := r.FormValue("email")
	phone_no := r.FormValue("phone_no")
	password := r.FormValue("password")
	fmt.Println(FullName)
	fmt.Println(email)
	fmt.Println(phone_no)
	fmt.Println(password)
}
