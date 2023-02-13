package configs

// import dependencies
import (
	"log"
	"os" // provides functions for interating w/ the os like access file, get file path

	"github.com/joho/godotenv"
)

// EnvMongoUri checks if the env variable is correctly loaded and return env variable.
// helper function that load the env variable using the godotenv library
func EnvMongoURI() string {

	// loads .env file in the current directory
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Getenv takes the name of the env variable("MONGODB_URI") and returns its associated. If the environment
	// variable specified does not exist on the system, the function returns an empty value.
<<<<<<< HEAD
	return os.Getenv("MONGODB_URI")
=======
	return os.Getenv("MONGOUDB_URI")
>>>>>>> dc0cdff345e54925f3c43d5c8eafab8165a08019
}
