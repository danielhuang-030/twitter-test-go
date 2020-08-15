package init

import (
	"app/model"
)

func init() {
	// connect DB
	model.ConnectDb()

	// connect RDB
	model.ConnectRdb()
}
