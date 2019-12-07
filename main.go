package main

import (
	"crypto/rsa"
	"net/http"
	"time"

	"log"
	"net/http/httputil"

	"github.com/rs/cors"
	"github.com/urfave/negroni"

	"github.com/techforward/auth-proxy/route"
)

var (
	verifyKey  *rsa.PublicKey
	signKey    *rsa.PrivateKey
	whiteHosts map[string]bool
)

var domainHost string = "lvh.me"

type middleware func(http.HandlerFunc) http.HandlerFunc

func main() {
	var routeProvider = route.New()
	whiteHosts = make(map[string]bool)
	whiteHosts["lvh.me"] = true
	whiteHosts["localhost"] = true

	director := func(request *http.Request) {
		destHost := routeProvider.GetDest(request.Host)
		log.Printf("%s to %s ", request.Host, destHost)
		// request.Host = destHost
		request.URL.Scheme = "http"
		request.URL.Host = destHost
		log.Printf("request.URL: %+v ", request.URL)
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"lvh.me", "localhost"},
		AllowCredentials: true,
		Debug:            false,
	})

	rp := &httputil.ReverseProxy{Director: director}

	n := negroni.New()

	n.Use(c)
	n.Use(negroni.HandlerFunc(AuthMiddleware))
	n.Use(negroni.HandlerFunc(FirewallMiddleware))
	n.UseHandler(rp)

	n.Use(negroni.NewLogger())

	server := http.Server{
		Addr:              ":80",
		Handler:           n,
		ReadHeaderTimeout: 20 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
