package auth

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type provider struct {
	Repository
}

var _ Repository = (*provider)(nil)

// New ...
func New() Repository {
	return &provider{}
}

func (p *provider) CreateToken() string {
	signBytes, err := ioutil.ReadFile("./private-key.pem")
	if err != nil {
		panic(err)
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic(err)
	}

	// create token
	token := jwt.New(jwt.SigningMethodRS256)

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "test"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		fmt.Println(err)
	}

	return tokenString
}

func (p *provider) VerifyToken(tokenString string) bool {
	verifyBytes, err := ioutil.ReadFile("./public-key.pem")

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)

	// verify token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, err := token.Method.(*jwt.SigningMethodRSA)
		if !err {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return verifyKey, nil

	})

	if err == nil && token.Valid {
		return true
	}
	return false
}
