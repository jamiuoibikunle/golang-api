package configs

import (
	"log"
	"os"
)

func EnvMongoURI() string {
	if err := os.Setenv("MONGODB_URI", "mongodb://localhost:27017"); err != nil {
		log.Fatal()
	}

	return os.Getenv("MONGODB_URI")
}
