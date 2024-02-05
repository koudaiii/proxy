package main

import (
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.Logger = log.New(log.Writer(), log.Prefix(), log.Flags())
	log.Println("Proxy server started at localhost:8080")
	// If you have any other proxy server, you can use only this proxy using this option.
	// ex. Client(curl --proxy localhost:8080 https://koudaiii.com) -> this proxy:8080 -> koudaiii.com:443
	proxy.ConnectDial = nil
	log.Fatal(http.ListenAndServe(":8080", proxy))
}
