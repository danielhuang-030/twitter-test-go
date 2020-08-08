package init

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	// load godotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
