package handler

import (
	"book/controller"
	"book/database"
	"book/model"
	"book/utils"
	"text/template"

	"net/http"

	"github.com/google/uuid"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// Ensure database connection
	db, err := database.DbIN()
	if err != nil {
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Check HTTP method
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusBadRequest)
		return
	}

	// Check URL path
	if r.URL.Path != "/register" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Create a new User object
	var newUser model.User
	newUser.ID = uuid.New()
	newUser.FullName = r.Form.Get("full_name")
	newUser.Email = r.Form.Get("email")
	newUser.PhoneNo = r.Form.Get("phone_no")
	newUser.Password = r.Form.Get("password")

	err = controller.InsertUser(db, newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// response := map[string]string{
	// 	"message": "User registered successfully",
	// }
	// json.NewEncoder(w).Encode(response)
	data := struct {
		FullName string
	}{
		FullName: newUser.FullName,
	}
	tmp, err := template.ParseFiles("./frontend/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmp.Execute(w, data)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// Ensure database connection
	db, err := database.DbIN()
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Check HTTP method
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusBadRequest)
		return
	}

	// Check URL path
	if r.URL.Path != "/login" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	token, fullNAme, err := controller.ValidateUser(db, w, email, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	_, err = utils.ValidateJWT(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	data := struct {
		FullName string
	}{
		FullName: fullNAme,
	}
	tmp, err := template.ParseFiles("./frontend/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmp.Execute(w, data)
	
	// json.NewEncoder(w).Encode(map[string]string{"token": token})

}
