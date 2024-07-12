package main

import (
	"book/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fileServer)
	http.HandleFunc("/register", handler.Register)
	
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
