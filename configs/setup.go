package configs

import (
	"context" // use to gather additional info about the env it being executed in.
	"fmt"     // format package allows to format strings, print, collect users' input.
	"log"     // provides basic logging features.
	"time"    // provides functionality for measuring and displaying time.

	"go.mongodb.org/mongo-driver/mongo" // Go driver lets you connect to and communicate w/ MongoDB cluster.
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* package: Context
It allow you to pass in a "context to your program". Context like a timeout
or deadline or a channel to indicate stop working and return. For instance
if you are doing a web request or running a system cmd, it is usually a good
idea to have a timout for production-grade systems. Because, if an API you
depend on is running slow, you would not want to back up

*/

// ConnectDB() have client to use the correct URI and check for errors. It return a ptr to a mongo.CLient
func ConnectDB() *mongo.Client {

	// mongo.NewClient() return *Client, error
	// options.Client creates a new ClientOption instance
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	// defined a timeout of 10 seconds when trying to connect.
	// "_" is a blank identifier, which ignores the returned value
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// client.Connect() initializes the Client
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	/* 	client.Ping() used to verify that the connection was created successfuly
	If the server is down, Ping will continue to reconnect until the client's
	timeout expires.
	*/
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client

}

// Client instance
var DB *mongo.Client = ConnectDB()

// GetCollection() retrieve and create collections on the database.
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golangAPI").Collection(collectionName)

	return collection
}
