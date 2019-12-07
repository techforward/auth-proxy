package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/techforward/auth-proxy/auth"
)

// AuthMiddleware 認証ミドルウェア
func AuthMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	log.Printf("AuthMiddleware")
	var authProvider auth.Repository = auth.New()

	if r.Host == domainHost && r.URL.String() == "/create" {
		token := authProvider.CreateToken()
		fmt.Fprint(w, token)
	}
	// var authProvider auth.Repository = auth.New()

	// // Authorization
	// if _, ok := r.Header["Proxy-Authorization"]; !ok {
	// 	tokenString := r.Header.Get("Proxy-Authorization")
	// 	isVerifyed := authProvider.VerifyToken(tokenString)
	// 	log.Printf("%s : %s", tokenString, isVerifyed)

	// 	if isVerifyed == false {
	// 		http.Error(w, "Token vailed.", http.StatusNotAcceptable)
	// 		return
	// 	}
	// }

	next(w, r)
}

// FirewallMiddleware ホワイトリスト
func FirewallMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	log.Printf("FirewallMiddleware")

	//Verify that request has an origin handler
	if r.Header.Get("Origin") == "" {
		http.Error(w, "Cross domain requests require Origin header", http.StatusBadRequest)
		return
	}
	origin := r.Header.Get("Origin")

	if _, ok := whiteHosts[origin]; ok {
		http.Error(w, "Failed Cros.", http.StatusBadRequest)
		return
	}

	next(w, r)
}
