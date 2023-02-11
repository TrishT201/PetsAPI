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
	"restAPI/configs"

	"github.com/gin-gonic/gin"
)

// Router responsibles for handling different route
func main() {
	// initialize a Gin router. THe Default configures Gin router w/ default middlewares (logger and recovery)
	router := gin.Default()

	// run database
	configs.ConnectDB()

	// Listen and serve HTTP request on localhost:6000. By default it serves on :8080 unless defined
	router.Run("localhost:6000")

}
