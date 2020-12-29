package util

import (
	"errors"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
  // ErrCreateToken for failures in creating Authentication Token
  ErrCreateToken = errors.New("Failed To Create Auth Token")
)

// CreateToken to create token
func CreateToken(userid uint64) (string, error) {
  var err error
  //Creating Access Token
  atClaims := jwt.MapClaims{}
  atClaims["authorized"] = true
  atClaims["user_id"] = userid
  atClaims["exp"] = time.Now().Add(time.Minute * 60 * 24 * 2).Unix()
  at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
  token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
  if err != nil {
    // ErrCreateToken.Error()
     return "", ErrCreateToken
  }
  return token, nil
}