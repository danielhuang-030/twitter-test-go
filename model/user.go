package model

import (
	"time"

	"gorm.io/gorm/clause"
)

type User struct {
	Model
	Name       string  `json:"name" gorm:"size:255;index"`
	Email      string  `json:"email" gorm:"size:255;uniqueIndex"`
	Password   string  `json:"-" gorm:"size:255"`
	Followers  []*User `json:"followers,omitempty" gorm:"many2many:user_followers;"`
	Followings []*User `json:"followings,omitempty" gorm:"many2many:user_followers;foreignKey:ID;joinForeignKey:FollowerID;References:ID;JoinReferences:UserID"`
	Posts      []Post  `json:"posts,omitempty"`
}

type UserFollower struct {
	UserID     uint      `json:"user_id" gorm:"unique:idx_user_follower"`
	FollowerID uint      `json:"follower_id" gorm:"unique:idx_user_follower"`
	CreatedAt  time.Time `json:"updated_at" gorm:"type:datetime"`
}

func CreateUser(data map[string]interface{}) (user User, err error) {
	user = User{
		Name:     data["name"].(string),
		Email:    data["email"].(string),
		Password: data["password"].(string),
	}
	if err = db.Omit(clause.Associations).Create(&user).Error; err != nil {
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

func FindUser(id uint) (user User, err error) {
	if err = db. /*.Preload(clause.Associations)*/ First(&user, id).Error; err != nil {
		return
	}

	return
}

func CreateFollower(user User, follower User) error {
	return db.Model(&user).Association("Followers").Append(&follower)
}

func DeleteFollower(user User, follower User) error {
	return db.Model(&user).Association("Followers").Delete(&follower)
}

func GetUserFollowers(id uint) (users []*User, err error) {
	user := User{}
	users = []*User{}
	if err = db.Preload("Followers").First(&user, id).Error; err != nil {
		return
	}
	if len(user.Followers) == 0 {
		return
	}
	users = user.Followers
	return
}

func GetUserFollowings(id uint) (users []*User, err error) {
	user := User{}
	users = []*User{}
	if err = db.Preload("Followings").First(&user, id).Error; err != nil {
		return
	}
	if len(user.Followings) == 0 {
		return
	}
	users = user.Followings
	return
}
