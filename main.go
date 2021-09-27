package main

import (
	"log"
	"net/http"

	"github.com/Lornzo/ReverseProxyServer/ReverseProxy"
	"github.com/Lornzo/ReverseProxyServer/configs"
)

func main() {
	configs.Load("/Users/a6288678/github.com/Lornzo/ReverseProxyServer/testdata/config.json")
	var reverse = ReverseProxy.New()
	log.Println("Reverse Proxy Server start working...")
	if configs.Get().UseSSL {
		http.ListenAndServeTLS(":443", configs.Get().CertFiles.CertFile, configs.Get().CertFiles.KeyFile, reverse)
	} else {
		http.ListenAndServe(":80", reverse)
	}

}
