// A rest api with mysql and mux router
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Book struct is for the book model
type Book struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Author    string `json:"author"`
	Published string `json:"published_at"`
}

// Init book var as slice
var book []Book

var db *sql.DB
var err error

// Retrieves all book data
func getBooks(w http.ResponseWriter, r *http.Request) {
	var book []Book
	w.Header().Set("Content-Type", "application/json")
	data, err := db.Query("SELECT id, name, author, published_at from book_info")
	if err != nil {
		panic(err)
	}
	defer data.Close()
	for data.Next() {
		var allBooks Book
		err = data.Scan(&allBooks.ID, &allBooks.Name, &allBooks.Author, &allBooks.Published)
		if err != nil {
			panic(err)
		}
		book = append(book, allBooks)
	}
	json.NewEncoder(w).Encode(book)
}

// Creates new book data
func makeBooks(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("INSERT INTO book_info (name, author, published_at) VALUES (?, ?, ?)")
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	name := keyVal["name"]
	author := keyVal["author"]
	publishedAt := keyVal["published_at"]

	_, err = stmt.Exec(name, author, publishedAt)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "New book data created")
}

// Updates an existing book data
func updateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE book_info SET name = ?, author = ?, published_at = ? WHERE id = ?")
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	newname := keyVal["name"]
	newauthor := keyVal["author"]
	newpublishedAt := keyVal["published_at"]

	_, err = stmt.Exec(newname, newauthor, newpublishedAt, params["id"])
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Book info with id:%s was updated", params["id"])
}

// Deletes an esisting book data
func deleteBooks(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	stmt, err := db.Prepare("DELETE FROM book_info WHERE id = ?")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Book info with id:%s was deleted", params["id"])
}

// Middleware function to authenticate required input
func required(input http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request received")
		input.ServeHTTP(w, r)
	}
}

func main() {
	// Establishing connection to database
	db, err = sql.Open("mysql", "admin:bo@tcp(127.0.0.1:3306)/bookdb")

	// Handle database connection error
	if err != nil {
		panic(err)
	} else if err = db.Ping(); err != nil {
		panic(err)
	}

	// Close database after func main has finished executing
	defer db.Close()
	fmt.Println("Database connection successful!")

	// Makes the book_info table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS book_info (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL, author VARCHAR(255) NOT NULL, published_at DATE)")
	if err != nil {
		panic(err)
	}

	// Initial sample data for database
	//book = append(book, Book{ID: "1", Name: "Silmarilion", Author: "JRR Tolkein", Published: "1977-08-15 08:00:00"})

	// Initialising the router
	router := mux.NewRouter()

	// Creating endpoints
	router.HandleFunc("/api/v1/books", getBooks).Methods("GET")
	router.HandleFunc("/api/v1/books", required(makeBooks)).Methods("POST")
	router.HandleFunc("/api/v1/books/{id}", updateBooks).Methods("PUT")
	router.HandleFunc("/api/v1/books/{id}", deleteBooks).Methods("DELETE")

	fmt.Println("server listening on port 8000")
	// Creates server on port 8000 with mux router
	log.Fatal(http.ListenAndServe(":8000", router))

}
