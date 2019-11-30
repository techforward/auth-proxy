package auth

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Provider ...
type Provider struct {
	r Repository
}

// NewProvider ...
func NewProvider(r Repository) *Provider {
	return &Provider{r}
}

func (p *Provider) create() string {
	signBytes, err := ioutil.ReadFile("./demo.rsa")
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

func (p *Provider) verify(tokenString string) bool {
	verifyBytes, err := ioutil.ReadFile("./demo.rsa.pub.pkcs8")
	if err != nil {
		panic(err)
	}
	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		panic(err)
	}

	// verify token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, err := token.Method.(*jwt.SigningMethodRSA)
		if !err {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		} else {
			return verifyKey, nil
		}
	})

	if err == nil && token.Valid {
		return true
	}
	return false
}
