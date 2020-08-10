package model

type User struct {
	Model
	Name     string `json:"name" gorm:"size:255;index"`
	Email    string `json:"email" gorm:"size:255;uniqueIndex"`
	Password string `json:"-" gorm:"size:255"`
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
	return
}

func GetUserByEmail(email string) (user User, err error) {
	if err = db.Where("email = ?", email).First(&user).Error; err != nil {
		return
	}

	return
}

func FindUser(id int) (user User, err error) {
	if err = db.First(&user, id).Error; err != nil {
		return
	}

	return
}
