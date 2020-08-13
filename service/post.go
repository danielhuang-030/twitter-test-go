package service

import (
	model "app/model"
)

func CreatePost(data map[string]interface{}, currentUser interface{}) (post model.Post, err error) {
	post, err = model.CreatePost(data, currentUser.(model.User))
	return
}
