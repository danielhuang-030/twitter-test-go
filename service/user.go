package service

import (
	model "app/model"
)

func GetUserInfo(id uint) (user model.User, err error) {
	user, err = model.FindUser(id)
	if err != nil {
		return
	}
	return
}
