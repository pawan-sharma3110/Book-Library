package controller

import (
	"book/model"
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func createUserTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        id UUID PRIMARY KEY,
        full_name VARCHAR(100) NOT NULL,
        email VARCHAR(100) NOT NULL UNIQUE,
        phone_no VARCHAR(15),
        password VARCHAR(100) NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT now()
    )`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		panic(err)

	}
	fmt.Println("Table created")
}
func InsertUser(db *sql.DB, user model.User) (err error) {
	createUserTable(db)

	insertUserSQL := `INSERT INTO users (id, full_name, email, phone_no, password, created_at)
    VALUES ($1, $2, $3, $4, $5, $6);`
	hasPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error while create passwod in hash: %v", err)
	}

	_, err = db.Exec(insertUserSQL, user.ID, user.FullName, user.Email, user.PhoneNo, hasPass, time.Now())
	if err != nil {
		return fmt.Errorf("unable to insert user: %v", err)
	}

	// fmt.Println("User inserted successfully")
	return nil
}
