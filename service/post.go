package service

import (
	model "app/model"
	"errors"
)

func CreatePost(data map[string]interface{}, currentUser interface{}) (post model.Post, err error) {
	post, err = model.CreatePost(data, currentUser.(model.User))
	return
}

func GetPost(id uint) (post model.Post, err error) {
	post, err = model.FindPost(id)
	if err != nil {
		return
	}
	return
}

func EditPost(id uint, data map[string]interface{}, currentUser interface{}) (post model.Post, err error) {
	post, err = model.FindPost(id)
	if err != nil {
		return
	}
	if err = checkPostPermission(post, currentUser); err != nil {
		return
	}
	post, err = model.UpdatePost(id, data)
	if err != nil {
		return
	}

	return
}

func DeletePost(id uint, currentUser interface{}) (err error) {
	post, err := model.FindPost(id)
	if err != nil {
		return
	}
	if err = checkPostPermission(post, currentUser); err != nil {
		return
	}
	return model.DeletePost(id)
}

func checkPostPermission(post model.Post, currentUser interface{}) error {
	if post.UserID != currentUser.(model.User).ID {
		return errors.New("You do not have enough permissions to perform this operation.")
	}
	return nil
}
