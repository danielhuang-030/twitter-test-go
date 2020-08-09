package model

type User struct {
	Model
	Name     string `json:"name" gorm:"size:255;index"`
	Email    string `json:"email" gorm:"size:255;uniqueIndex"`
	Password string `json:"password" gorm:"size:255"`
	Token    string `json:"token" gorm:"size:255"`
}

func CreateUser(data map[string]interface{}) User {
	user := User{
		Name:     data["name"].(string),
		Email:    data["email"].(string),
		Password: data["password"].(string),
	}
	db.Create(&user)

	return user
}
