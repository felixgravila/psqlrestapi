package main

// Author @felixgravila

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/felixgravila/psqlrestapi/dao"
	"github.com/felixgravila/psqlrestapi/model"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"strconv"
)

var connDetails = dao.ConnectionDetails{
	Host:     "localhost",
	Port:     5432,
	User:     "postgres",
	Password: "",
	Dbname:   "postgres",
}

var books []model.Book
var db *sql.DB

func connectToDb() {
	devmachine := os.Getenv("DEVMACHINE")
	if len(devmachine) == 0 {
		connDetails.Host = "postgres"
	}
	fmt.Println("PSQL using " + connDetails.Host)
	var err error
	db, err = dao.GetDB(connDetails)
	if err != nil {
		panic(err)
	}
}

func main() {
	r := mux.NewRouter()

	connectToDb()
	defer db.Close()

	r.HandleFunc("/authors", getAuthors).Methods("GET")
	r.HandleFunc("/authors/{id}", getAuthor).Methods("GET")
	r.HandleFunc("/authors/{firstname}/{lastname}", createAuthor).Methods("POST")
	r.HandleFunc("/authors", updateAuthor).Methods("PUT")
	r.HandleFunc("/authors/{id}", deleteAuthor).Methods("DELETE")
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books/{isbn}/{title}/{author_id}", createBook).Methods("POST")
	r.HandleFunc("/books", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// REST

func getAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	authors, err := dao.GetAuthors(db)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(authors)
	}
}

func getAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	val, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Can't parse id as int")
		return
	}
	author, err := dao.GetAuthor(db, val)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(author)
	}
}

func createAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	na, err := dao.AddAuthor(db, vars["firstname"], vars["lastname"])
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(na)
	}
}
func deleteAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	val, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Can't parse id as int")
		return
	}

	err = dao.DeleteAllBooksByAuthor(db, val)
	if err != nil {
		json.NewEncoder(w).Encode(fmt.Sprintf("Problem deleting books: %v", err))
		return
	}
	err = dao.DeleteAuthor(db, val)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode("OK")
	}
}
func updateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var a model.Author
	err := decoder.Decode(&a)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	err = dao.UpdateAuthor(db, &a)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(a)
	}

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books, err := dao.GetBooks(db)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(books)
	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	val, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Can't parse id as int")
		return
	}
	book, err := dao.GetBook(db, val)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(book)
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	aid, err := strconv.Atoi(vars["author_id"])
	if err != nil {
		json.NewEncoder(w).Encode("Cannot convert author_id to int")
		return
	}
	auth, err := dao.GetAuthor(db, aid)
	if err != nil {
		json.NewEncoder(w).Encode("Cannot find author")
		return
	}
	nb, err := dao.AddBook(db,
		vars["isbn"],
		vars["title"],
		auth,
	)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(nb)
	}
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	val, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Can't parse id as int")
		return
	}
	err = dao.DeleteBook(db, val)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode("OK")
	}
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var b model.Book
	err := decoder.Decode(&b)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	err = dao.UpdateBook(db, &b)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(b)
	}

}
