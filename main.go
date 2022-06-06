package main

import (
	"log"
	"net/http"
	"webapp/controllers"
)

func main() {
	http.HandleFunc("/books", controllers.Controller)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
