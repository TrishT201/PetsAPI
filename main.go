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

import "github.com/gin-gonic/gin"

// initalize a gin router
func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello from Gin-gonic & mongoDB",
		})
	})

	router.Run("localhost:6000")

}

type pet struct {
	ID    string  `json:"id"`
	PName string  `json:"pname"`
	DOB   float32 `json:"dob"`
	Owner string  `json:"owner"`
	Kind  string  `json:"kind"`
	Size  int     `json:"size"`
	Toy   string  `json:"toy"`
}
