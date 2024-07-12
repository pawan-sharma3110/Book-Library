package handler

import (
	"book/controller"
	"book/database"
	"book/model"
	"book/utils"
	"net/http"

	"github.com/gofrs/uuid"
)

func Register(w http.ResponseWriter, r *http.Request) {
	db, err := database.DbIN()
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
	}

	if r.Method != "POST" {
		http.Error(w, "Only post methord allowed", http.StatusBadRequest)
		return
	}
	if r.URL.Path != "/register" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	var newUser model.User
	if err := utils.ParseJson(w, r, newUser); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
	}
	id, err := uuid.NewV4()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	newUser.ID = id
	newUser.FullName = r.FormValue("full_name")
	newUser.Email = r.FormValue("email")
	newUser.PhoneNo = r.FormValue("phone_no")
	newUser.Password = r.FormValue("password")
	err = controller.InsertUser(db, newUser)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	utils.WriteJson(w, http.StatusOK, "user created")
}
