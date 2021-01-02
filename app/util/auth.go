package util

import (
	"errors"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/revel/revel"
	"github.com/tidwall/gjson"
)

var (
  // ErrCreateToken for failures in creating Authentication Token
  ErrCreateToken = errors.New("Failed To Create Auth Token")
  // ErrGettingTokenClaims for failure in retrieving claims from Authentication Token
  ErrGettingTokenClaims = errors.New("Failed to retrieve JWT Claims")
)

// CreateToken to create token
func CreateToken(userid uint64) (string, error) {
  var err error
  //Creating Access Token
  atClaims := jwt.MapClaims{}
  atClaims["authorized"] = true
  atClaims["user_id"] = userid
  // atClaims["exp"] = time.Now().Add(time.Minute * 60 * 24 * 2).Unix()
  at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
  token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
  if err != nil {
    // ErrCreateToken.Error()
     return "", ErrCreateToken
  }
  return token, nil
}

// ValidateToken is used to validate JWT Token
func ValidateToken(tokenString string) error {
  _, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    // return nil, nil
    // return struct{ Key []byte }{ Key: []byte(os.Getenv("ACCESS_SECRET")) }, nil
    return []byte(os.Getenv("ACCESS_SECRET")), nil
  })
  if err != nil {
    return errors.New("Failed to authenticate")
  }

  return nil
}

// GetTokenClaims is used to retrieve claims from JWT Token
func GetTokenClaims(tokenString string) (jwt.MapClaims, error) {
  // tokenString := "<YOUR TOKEN STRING>"    
  claims := jwt.MapClaims{}
  _, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
    return []byte(os.Getenv("ACCESS_SECRET")), nil
  })
  // ... error handling
  if err != nil {
    return claims, ErrGettingTokenClaims
  }
// do something with decoded claims
  // for key, val := range claims {
  //   fmt.Printf("Key: %v, value: %v\n", key, val)
  // }
  return claims, nil
}

// AuthInterceptor function provides authentication interception
func AuthInterceptor(redirectionRouteFunc func() revel.Result) func(c *revel.Controller) revel.Result {
  return func(c *revel.Controller) revel.Result {
    body := GetBodyString(c)
  if error := ValidateToken(gjson.Get(body, "data.token").String()); error != nil {
    return c.Redirect(redirectionRouteFunc)
  }
  return nil
  }
}