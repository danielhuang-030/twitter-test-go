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
	return model.CreateFollower(user, currentUser.(model.User))
}

func DeleteFollower(id uint, currentUser interface{}) (err error) {
	var user model.User
	user, err = model.FindUser(id)
	if err != nil {
		return
	}
	return model.DeleteFollower(user, currentUser.(model.User))
}
