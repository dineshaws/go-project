package common

import (
	"errors"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

// AUTHSECRET is jwt secret key
const (
	AUTHSECRET string = "mySecret"
)

// GenerateToken function to generate a new token
func GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})
	tokenStirng, err := token.SignedString([]byte(AUTHSECRET))
	return tokenStirng, err
}

// VerifyToken function is used to verify token
func VerifyToken(tokenString string) (string, error) {
	var username string
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(AUTHSECRET), nil
	})

	if err != nil {
		return username, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username = claims["username"].(string)
	} else {
		err = errors.New("Token is not valied")
	}
	return username, err
}
