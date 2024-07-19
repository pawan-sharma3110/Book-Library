package controller

import (
	"book/model"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("your_secret_key")

func createUserTable(db *sql.DB) {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		full_name TEXT,
		email TEXT UNIQUE,
		phone_no TEXT,
		password TEXT,
		created_at TIMESTAMP
	);
	`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("unable to create table: %v", err)
	}
}

func InsertUser(db *sql.DB, user model.User) error {
	// Create user table if not exists
	createUserTable(db)

	insertUserSQL := `
		INSERT INTO users (id, full_name, email, phone_no, password, created_at)
		VALUES ($1, $2, $3, $4, $5, $6);
	`
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error while creating password hash: %v", err)
	}

	_, err = db.Exec(insertUserSQL, user.ID, user.FullName, user.Email, user.PhoneNo, hashedPassword, time.Now())
	if err != nil {
		return fmt.Errorf("provided email already register")
	}

	return nil
}

func setCookies(w http.ResponseWriter, userID uuid.UUID, token, email string) {
	expirationTime := time.Now().Add(24 * time.Hour)

	http.SetCookie(w, &http.Cookie{
		Name:     "user_id",
		Value:    userID.String(),
		Expires:  expirationTime,
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  expirationTime,
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "email",
		Value:    email,
		Expires:  expirationTime,
		HttpOnly: true,
	})
}
func ValidateUser(db *sql.DB, w http.ResponseWriter, email string, password string) (tokenString string, err error) {
	type Claims struct {
		Email string    `json:"email"`
		ID    uuid.UUID `json:"id"`
		jwt.StandardClaims
	}
	// Fetch user from the database
	var user model.User
	err = db.QueryRow("SELECT id, full_name, email, phone_no, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.FullName, &user.Email, &user.PhoneNo, &user.Password)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return "", err
	}
	// Compare hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return "", err
	}
	// Generate JWT
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: user.Email,
		ID:    user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err = token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	setCookies(w, user.ID, tokenString, user.Email)
	return tokenString, nil
}
