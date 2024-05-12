package apiendpoints

import (
	// "context"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func Test(name string) string {
	return ("hello" + name)
}

func getBooks(c *gin.Context) []Book {
	db, err := sql.Open("mysql", "admin:password@tcp(parker-database.cfhfkqv5cjrl.us-east-1.rds.amazonaws.com:3306)/book_schema")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()

	query := "SELECT * FROM books"

	queryResult, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}

	var books []Book

	for queryResult.Next() {
		var bk Book
		if err := queryResult.Scan(&bk.ID, &bk.Title); err != nil {
			fmt.Println(books, err)
		}
		books = append(books, bk)

	}

	return books
}
