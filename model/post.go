package model

type Post struct {
	Model
	UserID  uint   `json:"user_id" gorm:"index"`
	Content string `json:"content" gorm:"type:text"`
}

func CreatePost(data map[string]interface{}, user User) (post Post, err error) {
	post = Post{
		Content: data["content"].(string),
	}

	if err = db.Model(&user).Association("Posts").Append(&post); err != nil {
		return
	}
	return
}

func FindPost(id uint) (post Post, err error) {
	if err = db.First(&post, id).Error; err != nil {
		return
	}

	return
}

func UpdatePost(id uint, data map[string]interface{}) (post Post, err error) {
	post, err = FindPost(id)
	if err != nil {
		return
	}
	if err = db.Model(&post).Updates(data).Error; err != nil {
		return
	}

	return
}

func DeletePost(id uint) (err error) {
	post, err := FindPost(id)
	if err != nil {
		return
	}
	return db.Delete(&post).Error
}
