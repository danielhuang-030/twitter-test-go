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

func GetUserFollowers(id uint) (users []*model.User, err error) {
	users, err = model.GetUserFollowers(id)
	if err != nil {
		return
	}

	return
}

func GetUserFollowings(id uint) (users []*model.User, err error) {
	users, err = model.GetUserFollowings(id)
	if err != nil {
		return
	}

	return
}
