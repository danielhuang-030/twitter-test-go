package model

import "errors"

type User struct {
	Model
	Name     string `json:"name" gorm:"size:255;index"`
	Email    string `json:"email" gorm:"size:255;uniqueIndex"`
	Password string `json:"-" gorm:"size:255"`
	Token    string `json:"token" gorm:"size:255"`
}

func CreateUser(data map[string]interface{}) (user User, err error) {
	user = User{
		Name:     data["name"].(string),
		Email:    data["email"].(string),
		Password: data["password"].(string),
	}
	if err = db.Create(&user).Error; err != nil {
		return
	}
	return user, err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	db.Where("email = ?", email).First(&user)
	if 0 == int(user.ID) {
		err = errors.New("The user is not exist")
	}

	return
}
