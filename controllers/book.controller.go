package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"webapp/models"
	s "webapp/services"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		read(w, r)
	case "POST":
		create(w, r)
	}
}

func read(w http.ResponseWriter, r *http.Request) {
	books, err := s.GetAllBooks()
	if err != nil {
		log.Fatal(err)
	}

	jsonResponse, err := json.Marshal(books)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func create(w http.ResponseWriter, r *http.Request) {
	// b := make([]byte, 256)
	// var body string
	// for {
	// 	n, err := r.Body.Read(b)
	// 	body = body + string(b[:n])
	// 	if err == io.EOF {
	// 		break
	// 	}
	// }

	var book models.Book
	decodeErr := json.NewDecoder(r.Body).Decode(&book)
	if decodeErr != nil {
		log.Fatal(decodeErr)
	}

	newBook, err := s.CreateBook(book)
	if err != nil {
		log.Fatal(err)
	}

	response, err := json.Marshal(newBook)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(response)
}
