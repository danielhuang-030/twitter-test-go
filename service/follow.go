package service

import (
	model "app/model"
)

func CreateFollower(id uint, currentUser interface{}) (err error) {
	var user model.User
	user, err = model.FindUser(id)
	if err != nil {
		return
	}
	return model.GreateFollower(user, currentUser.(model.User))
}
