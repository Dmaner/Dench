package database

import (
	"fmt"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

func test() {

	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
	})
	if err != nil {
		fmt.Println("test")
	}
	c, err := driver.NewClient(driver.ClientConfig{
		Connection: conn,
	})
	if err != nil {
		fmt.Println("test")
	}
	// Open "examples_books" database
	db, err := c.Database(nil, "examples_books")
	if err != nil {
		fmt.Println("test")
	}

	// Open "books" collection
	col, err := db.Collection(nil, "books")
	if err != nil {
		fmt.Println("test")
	}

	fmt.Printf("Created document in collection '%s' in database '%s'\n", col.Name(), db.Name())
}
