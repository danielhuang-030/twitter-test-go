package init

import (
	"github.com/joho/godotenv"
)

func init() {
	// set timezone
	godotenv.Load()
}
