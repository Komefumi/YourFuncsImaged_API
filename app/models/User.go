package models

import (
	"errors"
	"fmt"

	util "github.com/Komefumi/YourFuncsImaged_API/app/util"

	b "golang.org/x/crypto/bcrypt"
)

var (
	NotFoundError = errors.New("User Not Found")
	PasswordNotMatchError = errors.New("Password did not match")
)

// User model to be exported and used
type User struct {
	ID uint `gorm:"primaryKey;column:id;autoIncrement"`
	Email string `gorm:"column:email"`
	PasswordHash string `gorm:"column:password_hash"`
}

// FindAuthUser by email and password
func FindAuthUser(email, password string) (*User, error) {
	var foundUser User;
	db := util.DBAccessorFunc()
	fmt.Println("Email is", email)
	// db.Where(&User{ Email: email }).First(&foundUser)
	// db.Raw("SELECT id, email, password_hash FROM users WHERE ID = 1", email).Scan(foundUser)
	db.Table("users").Select("id", "email", "password_hash").Where("email = ?", email).Scan(&foundUser)
	fmt.Println("Found ", foundUser)
	if (&foundUser == nil) {
		return nil, NotFoundError
	}
	fmt.Println(password)
	passwordVerified := foundUser.VerifyPassword(password)
	if (passwordVerified == false) {
		return nil, PasswordNotMatchError
	}

	return &foundUser, nil
}

// GetUser by ID
func GetUser(id uint) User {
	var foundUser User;
	db := util.DBAccessorFunc()
	db.Raw("SELECT id, email, password_hash FROM users WHERE id = ?", id).Scan(&foundUser)
	return foundUser
}

// SetPassword used for setting password_hash of user
func(user *User) SetPassword(password string) {
	passwordArray := []byte(password)
	hashedPassword, err := b.GenerateFromPassword(passwordArray, b.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.PasswordHash = string(hashedPassword)
	return
}

// SetEmail used for setting email of user
func(user *User) SetEmail(email string) {
	user.Email = email
}

// VerifyPassword checks if the plaintext is the password
func(user *User) VerifyPassword(plaintext string) bool {
	err := b.CompareHashAndPassword([]byte(user.PasswordHash), []byte(plaintext))
	if err != nil {
		return false
	}
	return true
}

