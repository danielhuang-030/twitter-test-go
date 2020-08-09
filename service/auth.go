package service

import (
	model "app/model"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(data map[string]interface{}) (user model.User, err error) {
	data["password"] = hashAndSalt(data["password"].(string))
	user, err = model.CreateUser(data)
	return
}

func Attempt(data map[string]interface{}) (user model.User, err error) {
	user, err = model.GetUserByEmail(data["email"].(string))
	if err != nil {
		return
	}

	if !comparePasswords(string(user.Password), data["password"].(string)) {
		err = errors.New("wrong password")
		return
	}

	// // set token
	// $tokenResult = $user->createToken(static::TOKEN_KEY);
	// $tokenResult->token->save();
	// $user->withAccessToken($tokenResult->accessToken);

	// // set user
	// Auth::setUser($user);

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
