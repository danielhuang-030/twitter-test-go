package init

import (
	"os"
)

func init() {
	// set timezone
	os.Setenv("TZ", os.Getenv("APP_TIMEZONE"))
}
