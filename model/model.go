package model

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}

var (
	db  *gorm.DB
	rdb *redis.Client
)

func ConnectDb() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// many to many
	db.SetupJoinTable(&User{}, "Followers", &UserFollower{})

	db.AutoMigrate(&User{}, &UserFollower{}, &Post{})
}

func GetDb() *gorm.DB {
	if db == nil {
		ConnectDb()
	}
	return db
}

func ConnectRdb() {
	opt, err := redis.ParseURL(fmt.Sprintf("redis://%s:%s/%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"), os.Getenv("REDIS_DB")))
	if err != nil {
		panic(err)
	}
	rdb = redis.NewClient(opt)
}

func GetRdb() *redis.Client {
	if rdb == nil {
		ConnectRdb()
	}
	return rdb
}
