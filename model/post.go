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
