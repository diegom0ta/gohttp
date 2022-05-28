package main

import (
	"fmt"
	"net/http"
	"context"
)

var srv *http.Server

func readUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func getClientIP(w http.ResponseWriter, req *http.Request) {
	if ip := readUserIP(req); ip != "" {
		fmt.Println(ip)
		srv.Shutdown(context.TODO())
	}
}

func main() {
	port := ":8090"
	
	srv = new(http.Server)
	srv.Addr = port

	http.HandleFunc("/myip", getClientIP)
	srv.ListenAndServe()
}	
