package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	pusher "github.com/pusher/pusher-http-go"
)

// Comment is a struct
type Comment struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Comment string `json:"comment"`
}

// We register the Pusher client
var client = pusher.Client{
	AppId:   "PUSHER_APP_ID",
	Key:     "PUSHER_APP_KEY",
	Secret:  "PUSHER_APP_SECRET",
	Cluster: "PUSHER_APP_CLUSTER",
	Secure:  true,
}

// CommentCollection is collection of Comments
type CommentCollection struct {
	Comments []Comment `json:"items"`
}

// GetComments retrieves data for the comments from the database
func GetComments(db *sql.DB) CommentCollection {

	sql := "SELECT * FROM comments"
	rows, err := db.Query(sql)

	// Exit if the SQL doesn't work for some reason
	if err != nil {
		panic(err)
	}

	// make sure to cleanup when the program exits
	defer rows.Close()

	result := CommentCollection{}

	for rows.Next() {
		comment := Comment{}
		err2 := rows.Scan(&comment.ID, &comment.Name, &comment.Email, &comment.Comment)

		// Exit if we get an error
		if err2 != nil {
			panic(err2)
		}
		result.Comments = append(result.Comments, comment)
	}
	return result
}

// PushComment updates the database using an index
func PushComment(db *sql.DB, name string, email string, comment string) (int64, error) {
	sql := "INSERT INTO comments(name, email, comment) VALUES(?, ?, ?)"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the ?'s in our prepared statement with 'name', 'email' and 'comment'
	result, err2 := stmt.Exec(name, email, comment)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	newComment := Comment{
		Name:    name,
		Email:   email,
		Comment: comment,
	}

	client.Trigger("comment-channel", "new-comment", newComment)

	return result.LastInsertId()
}
