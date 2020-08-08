package model

type User struct {
	Model
	Name     string `gorm:"size:255;index"`
	Email    string `gorm:"size:255;uniqueIndex"`
	Password string `gorm:"size:255"`
	Token    string
}
