package init

import (
	model "app/model"
)

func init() {
	model.ConnectDb()
}
