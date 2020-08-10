package service

import (
	model "app/model"
)

func GetUserInfo(id int) (user model.User, err error) {
	user, err = model.FindUser(id)
	if err != nil {
		return
	}
	return
}
