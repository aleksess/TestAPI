package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// structs

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstn"`
	Lastname  string `json:"lastn"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(10000000) + 3)
	books = append(books, book)

	json.NewEncoder(w).Encode(book)
}

func main() {
	router := mux.NewRouter()

	books = append(books, Book{"1", "Lalka", &Author{"Boleslaw", "Prus"}})
	books = append(books, Book{"2", "Sachem", &Author{"Henryk", "Sienkiewicz"}})

	//handlers

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	/*router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")*/
	fmt.Println("App listening on port 8000!")

	http.ListenAndServe(":8000", router)

}
