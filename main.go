package main

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"log"
	"net/http"
	"net/http/httputil"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)
token := &auth.NewProvider{}

func main() {
	// provier := &NewProvider{}

	director := func(request *http.Request) {
		// Authorization
		if _, ok := request.Header["Proxy-Authorization"]; !ok {

			tokenString := request.Header.Get("Proxy-Authorization")

			isVerifyed := token.verify(tokenString)

			if token.Valid {

			}
		}

		request.URL.Scheme = "http"
		request.URL.Host = ":3000"

	}

	rp := &httputil.ReverseProxy{Director: director}
	server := http.Server{
		Addr:    ":80",
		Handler: rp,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}

// LoginHandler : JWTの発行
func LoginHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//今回、username,password共にtestを使っていますが、皆様が実際に行う場合はbodyから読み込んだり、DBから読み込んだりして下さい。
	username := "test"
	password := "test"

	signBytes, err := ioutil.ReadFile("./demo.rsa")
	if err != nil {
		panic(err)
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic(err)
	}

	if username == "test" && password == "test" {
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
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(tokenString))
	}
}
