/*
router: Takes the request and decides which controller/controller

	methods will handle the request. The controller accepts the
	request and handles it.
*/
package routes

import (
	"restAPI/controllers"

	"github.com/gin-gonic/gin"
)

// CRUD routes to call the controllers
func Route(router *gin.Engine) {
	// ALL routers related to users come here
	router.POST("/pet", controllers.CreatePet())
	router.GET("/pet/:petId", controllers.GetPet())
	router.PUT("/pet/:petId", controllers.EditPet())
	router.DELETE("/pet/:petId", controllers.DelPet())
	router.GET("/pets", controllers.GetAllPet())
}
