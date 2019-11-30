package main

import (
	"crypto/rsa"
	"net/http"

	"log"
	"net/http"
	"net/http/httputil"

	"github.com/rs/cors"

	"github.com/techforward/auth-proxy/auth"
)

var (
	verifyKey  *rsa.PublicKey
	signKey    *rsa.PrivateKey
	whiteHosts map[string]bool
)

func main() {
	var token = auth.NewProvider()
	var domainHost = "lvh.me"
	whiteHosts = make(map[string]bool)

	director := func(request *http.Request) {
		// Authorization
		if _, ok := request.Header["Proxy-Authorization"]; !ok {

			tokenString := request.Header.Get("Proxy-Authorization")

			isVerifyed := token.verifyToken(tokenString)

			if token.Valid {

			}
		}

		request.URL.Scheme = "http"
		request.URL.Host = ":3000"

	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"lvh.me"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	rp := &httputil.ReverseProxy{Director: director}
	server := http.Server{
		Addr:    ":80",
		Handler: c.Handler(rp),
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
