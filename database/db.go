package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

type Book struct {
	ID     int    `json:"id" xml:"id"`
	Title  string `json:"title" xml:"title"`
	Author string `json:"author" xml:"author"`
	Year   string `json:"year" xml:"year"`
}

var books []Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(mes string, err error) {
	if err != nil {
		fmt.Println(mes)
		log.Fatal("", err)
		fmt.Println(mes)
	}
}
func startDatabase() {
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal("this is in the LoadDatabase function", err)
	db, err = sql.Open("postgres", pgUrl)
	err = db.Ping()
}

//LoadDatabase loads the information in the database
func LoadDatabase() []Book {
	var book Book
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal("this is in the LoadDatabase function", err)
	db, err = sql.Open("postgres", pgUrl)
	err = db.Ping()
	logFatal("this is in the LoadDatabase function", err)

	rows, err := db.Query("select * from books")

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal("this is in the LoadDatabase function", err)
		books = append(books, book)
	}

	return books
}

//LoadSingleRow loads a single row from the database
func LoadSingleRow(id string) Book {
	startDatabase()

	var book Book
	fmt.Println(id)
	rows := db.QueryRow("select * from books where id=$1", id)
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal("LoadSingleRow", err)
	logFatal("This is in the LoadSingleRow", err)
	// books = append(books, book)
	// fmt.Println(books)
	return book
}

// AddDetails adds details from the front end to the database
func AddDetails(book Book) int {
	var bookID int

	startDatabase()
	err := db.QueryRow("insert into books (title,author,year) values($1,$2,$3)RETURNING id;", book.Title, book.Author, book.Year).Scan(&bookID)
	logFatal("this is in the adddetails func", err)
	return bookID

}

//UpdateTables updates the values in the database
func UpdateTables(book Book) int64 {
	fmt.Println(&book.Title, &book.Author, &book.Year, &book.ID)

	result, _ := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id;",
		book.Title, book.Author, book.Year, book.ID)
	// logFatal("this is in the update func", err)
	rowsAffected, er := result.RowsAffected()
	logFatal("this is in the updattables func", er)
	return rowsAffected
}

//RemoveDetails remove a detail from the table
func RemoveDetails(id string) int64 {
	result, err := db.Exec("delete from books where id = $1", id)
	logFatal("this is in the remove func", err)

	rowsDeleted, err := result.RowsAffected()
	logFatal("this is from the second err in the remove func", err)

	return rowsDeleted

}
