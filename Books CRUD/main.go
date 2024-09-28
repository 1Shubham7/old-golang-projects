package main

import
(
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

type Book struct {
	ID string `json:"id"` // these are called struct tags
	Price string `json:"paise kitne lagenge"`
	Title string `json:"title"`
	Author *Author `json:"-"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var books []Book

func main() {

	// adding some default books 
	books = append(books, Book{ID:"1", Price:"2323", Title:"Meri book", Author: &Author{Firstname:"Shubham", Lastname:"Singh"}})
	books = append(books, Book{ID:"1", Price:"1000", Title:"Acchi book", Author: &Author{Firstname:"Shubham", Lastname:"Singh"}})
	books = append(books, Book{ID:"1", Price:"1", Title:"Best book", Author: &Author{Firstname:"Shubham", Lastname:"Singh"}})

	r := mux.NewRouter()
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	fmt.Printf("Starting the server at port 1111. Go check 'http://localhost:1111/books'\n")
	log.Fatal(http.ListenAndServe(":1111",r))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books) 
}

func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)  // this means instead of book[:index] have book[index+1] therefore that entry is removed
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _,item := range books {
		if item.ID == params["id"]{
		json.NewEncoder(w).Encode(item) // simply return it
		return
		}
	}
}

func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000000)) //strconv converts it into string
	books =  append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json") //set the json content type
	params := mux.Vars(r)
	for index,item := range books{
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...) // this means instead of books[:index] have books[index+1] therefore that entry is removed
			var book Book 
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"] //we want the neww one to go to id of old one
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return

		}
	}
}