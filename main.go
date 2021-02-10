package main

import (
	"log"
	"net/http"

	"github.com/draiz/Learning/Books-list/handlers"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

// func logFatal(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// var db *sql.DB

func main() {
	//Make sure you setup the ELEPHANTSQL_URL to be a uri, e.g. 'postgres://user:pass@host/db?options'

	// pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	// logFatal(err)
	// db, err = sql.Open("postgres", pgUrl)
	// err = db.Ping()
	// logFatal(err)
	router := mux.NewRouter()

	router.HandleFunc("/books", handlers.GetBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", handlers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", handlers.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books", handlers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", handlers.RemoveBook).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8000", router))
}
