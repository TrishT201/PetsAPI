package main

/*	CRUD Operation
	POST - Create
	GET - Read
	PUT	- Update
	DELETE - Delete

	Status Code:
	200 lvl - success
	400 lvl - something wrong w the request
	500 lvl - something wrong at the server

	gin/gin gonic - HTTP web framework

	mongoDB - a document-based db management program used as an alternative to relational db.
*/

// import dependencies
import (
	"github.com/gin-gonic/gin"
)

func main() {
	// initialize a Gin router. THe Default configures Gin router w/ default middlewares (logger and recovery)
	router := gin.Default()

	// .Get is used to defined a handler for a GET request. It takes in route(/) and one or more route handlers as parameters.
	router.GET("/", func(c *gin.Context) {

		// JSON(code int, obj interface{}) renders the given obj as json
		c.JSON(200, gin.H{
			"data": "Hello from Gin-gonic & mongoDB",
		})
	})

	// Listen and serve HTTP request on localhost:6000. By default it serves on :8080 unless defined
	router.Run("localhost:6000")

}

type pet struct {
	ID    string  `json:"id"`
	PName string  `json:"pname"`
	DOB   float64 `json:"dob"`
	Owner string  `json:"owner"`
	Kind  string  `json:"kind"`
	Size  int     `json:"size"`
	Toy   string  `json:"toy"`
}
