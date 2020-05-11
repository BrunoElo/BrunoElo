package main

import (
	"log"
	"strconv"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"encoding/json"
)

// Book struct is for the book model
type Book struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Author string `json:"author"`
	Published string `json:"published_at"`
}

// Init book var as slice
var book []Book

// Retrieves all book data
func getBooks(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// Creates new book data
func makeBooks(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	var newBook Book
	json.NewDecoder(r.Body).Decode(&newBook)
	newBook.ID = strconv.Itoa(len(book) + 1)
	book = append(book, newBook)
	json.NewEncoder(w).Encode(newBook)
}

// Updates an existing book data
func updateBooks(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range book {
		if item.ID == params["id"] {
			// Using append() to remove the old book
			book = append(book[:i], book[i+1:]...)
			var newBook Book
			json.NewDecoder(r.Body).Decode(&newBook)
			newBook.ID = params["id"]
			book = append(book, newBook)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// Deletes an esisting book data
func deleteBooks(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range book {
		if item.ID == params["id"] {
		book = append(book[:i], book[i+1:]...)
		break
		}
	}
	json.NewEncoder(w).Encode(book)
}

// Middleware function to authenticate required input
func required(input http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request received")
		input.ServeHTTP(w, r)
	}
}

func main()  {
	// Sample data
	book = append(book, Book{ID:"1", Name: "Silmarilion", Author: "JRR Tolkein", Published: "1977-08-15T08:00:00"})

	// Initialising the router
	router := mux.NewRouter()

	// Creating endpoints
	router.HandleFunc("/api/v1/books", required(getBooks)).Methods("GET")
	router.HandleFunc("/api/v1/books", required(makeBooks)).Methods("POST")
	router.HandleFunc("/api/v1/books/{id}", required(updateBooks)).Methods("PUT")
	router.HandleFunc("/api/v1/books/{id}", required(deleteBooks)).Methods("DELETE")

	fmt.Println("server listening on port 8000")
	// Creates server on port 8000 with mux router
	log.Fatal(http.ListenAndServe(":8000", router))
	
}