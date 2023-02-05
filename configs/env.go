package configs

// import dependencies
import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// EnvMongoUri checks if the env variable is correctly loaded and return env variable
// helper function to load the env variable using the godotenv library
func EnvMongoURI() string {

	// loads .env file in the current directory
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Getenv takes the name of the env variable("MONGOURI") and returns its associated. If the environment
	// variable specified does not exist on the system, the function returns an empty value.
	return os.Getenv("MONGOURI")
}
