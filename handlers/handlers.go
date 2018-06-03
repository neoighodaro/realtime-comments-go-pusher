package handlers

import (
	"database/sql"
	"go-realtime-comments/models"
	"net/http"

	"github.com/labstack/echo"
)

// H can map strings as keys and anything as values
type H map[string]interface{}

//GetComments handles the HTTP request that hits the /comments endpoint
func GetComments(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Fetch comments using our new model
		return c.JSON(http.StatusOK, models.GetComments(db))
	}
}

// PushComment handles the incoming HTTP requests
func PushComment(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new comment
		var comment models.Comment
		// Map imcoming JSON body to the new Comment
		c.Bind(&comment)
		// Add a comment using our new model
		id, err := models.PushComment(db, comment.Name, comment.Email, comment.Comment)
		// Return a JSON response if successful
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
			// Handle any errors
		} else {
			return err
		}
	}
}
