package service

import (
	"app/middleware"
	"app/model"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(data map[string]interface{}) (user model.User, err error) {
	data["password"] = hashAndSalt(data["password"].(string))
	user, err = model.CreateUser(data)
	return
}

func Attempt(data map[string]interface{}) (user model.User, token string, err error) {
	user, err = model.GetUserByEmail(data["email"].(string))
	if err != nil {
		return
	}

	if !comparePasswords(string(user.Password), data["password"].(string)) {
		err = errors.New("wrong password")
		return
	}

	// set token
	token, err = generateToken(user.ID, data["email"].(string))
	if err != nil {
		return
	}

	return
}

func Logout(token string, expireAt int64) (err error) {
	err = model.CreateBlacklistToken(token, expireAt)
	if err != nil {
		return
	}
	return
}

func hashAndSalt(pwd string) string {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	bytePlainPwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// @from https://github.com/eddycjy/go-gin-example/blob/master/pkg/util/jwt.go
func generateToken(id uint, email string) (string, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	nowTime := time.Now()
	envExpireMin, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE_MIN"))
	expireTime := nowTime.Add(time.Duration(envExpireMin) * time.Minute)

	claims := middleware.Claims{
		id,
		encodeMD5(email),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// @from https://github.com/eddycjy/go-gin-example/blob/master/pkg/util/md5.go
func encodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
