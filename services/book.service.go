package services

import (
	"fmt"
	"log"
	infra "webapp/infrastructure"
	m "webapp/models"
)

func GetAllBooks() (books []m.Book, err error) {
	db := infra.GetInstance()
	rows, err := db.Query("SELECT * FROM book")

	if err != nil {
		fmt.Printf("Error queriying for books %v", err)
	}

	for rows.Next() {
		var book m.Book
		if err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Year); err != nil {
			fmt.Printf("Error queriying for books %v", err)
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("Error queriying for books %v", err)
	}
	rows.Close()
	return
}

func CreateBook(book m.Book) (newBook m.Book, err error) {
	db := infra.GetInstance()
	row := db.QueryRow(`
	INSERT INTO book (title, author, year) 
	VALUES ($1, $2, $3) 
	RETURNING id, title, author, year`, book.Title, book.Author, book.Year)

	if err := row.Scan(&newBook.Id, &newBook.Title, &newBook.Author, &newBook.Year); err != nil {
		log.Fatal(err)
	}
	return
}
