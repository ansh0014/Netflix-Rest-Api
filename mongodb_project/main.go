package main

import (
	"fmt"
	"log"
	 "mongodb_project/router"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is starting at: http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}