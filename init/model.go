package init

import (
	model "app/model"
)

func init() {
	// connect DB
	model.ConnectDb()
}
