/*
	 controller: responsible for handling the HTTP requests
		coming given by the router. It's an interface that allows
		Model to interact w/ a certain medium AKA HTTP request.
*/
package controllers

import (
	"context"
	"net/http"
	"restAPI/configs"
	"restAPI/models"
	"restAPI/responses"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var petCollection *mongo.Collection = configs.GetCollection(configs.DB, "pets")

// CreatePet() returns a Gin-gonic handler function
func CreatePet() gin.HandlerFunc {
	return func(c *gin.Context) {

		// defined a timeout when inserting pet into the document
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var pet models.Pet
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&pet); err != nil {
			c.JSON(http.StatusBadRequest, responses.PetResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// create a newPet variable
		newPet := models.Pet{
			ID:    primitive.NewObjectID(),
			PName: pet.PName,
			DOB:   pet.DOB,
			Owner: pet.Owner,
			Kind:  pet.Kind,
			Size:  pet.Size,
			Toy:   pet.Toy,
		}

		// insert newPet into the collection
		result, err := petCollection.InsertOne(ctx, newPet)

		// check if the insert was successful
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PetResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// the insert was successful
		c.JSON(http.StatusCreated, responses.PetResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetPet() gin.HandlerFunc {
	return func(c *gin.Context) {

		// defined a timeout when finding a pet in the document
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		// get pet's id from url param
		petId := c.Param("petId")
		var pet models.Pet
		defer cancel()

		// convert petID from a string to a primitive.ObjectID type
		objId, _ := primitive.ObjectIDFromHex(petId)

		// .FindOne() search for the pet by objID
		err := petCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&pet)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PetResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.PetResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": pet}})

	}
}

func EditPet() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		petId := c.Param("petId")
		var pet models.Pet
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(petId)

		// validate the request body
		if err := c.BindJSON(&pet); err != nil {
			c.JSON(http.StatusBadRequest, responses.PetResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		update := bson.M{"name": pet.PName, "owner": pet.Owner, "kind": pet.Kind, "size": pet.Size, "toy": pet.Toy}
		result, err := petCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PetResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// get updated pet details
		var updatedPet models.Pet
		if result.MatchedCount == 1 {
			err := petCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedPet)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.PetResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}
		}

		c.JSON(http.StatusOK, responses.PetResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedPet}})
	}
}

func DelPet() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		petId := c.Param("petId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(petId)

		result, err := petCollection.DeleteOne(ctx, bson.M{"id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PetResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusInternalServerError, responses.PetResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": "Pet with specified ID not found!"}})
			return
		}

		c.JSON(http.StatusOK, responses.PetResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted"}})

	}
}

func GetAllPet() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var pets []models.Pet
		defer cancel()

		results, err := petCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.PetResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		// reading from db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singlePet models.Pet
			if err = results.Decode(&singlePet); err != nil {
				c.JSON(http.StatusInternalServerError, responses.PetResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}
			pets = append(pets, singlePet)
		}

		c.JSON(http.StatusOK, responses.PetResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": pets}})
	}

}
