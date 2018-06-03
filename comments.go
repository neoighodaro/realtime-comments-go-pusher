package main

import (
	"database/sql"
	"go-realtime-comments/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// Here we check for any db errors then exit
	if err != nil {
		panic(err)
	}

	// If we don't get any errors but somehow still don't get a db connection
	// we exit as well
	if db == nil {
		panic("db nil")
	}
	return db
}

// Here we create a function to migrate the database and insert the first rows for the votes
func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS comments(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR NOT NULL,
			email VARCHAR NOT NULL,
			comment VARCHAR NOT NULL
	);
   `
	_, err := db.Exec(sql)

	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize the database
	db := initDB("storage.db")
	migrate(db)

	// Define the HTTP routes
	e.File("/", "public/index.html")
	e.GET("/comments", handlers.GetComments(db))
	e.POST("/comment", handlers.PushComment(db))

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}
