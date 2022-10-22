package main

import (
	"context"
	"log"
	"net/http"
)

const port = ":8090"

var srv = new(http.Server)

func readClientIp(r *http.Request) string {
	ipAddr := r.Header.Get("X-Real-Ip")
	if ipAddr == "" {
		ipAddr = r.Header.Get("X-Forwarded-For")
	}
	if ipAddr == "" {
		ipAddr = r.RemoteAddr
	}
	return ipAddr
}

func getClientIp(w http.ResponseWriter, req *http.Request) {
	defer srv.Shutdown(context.TODO())
	ip := readClientIp(req)
	log.Printf("Client IP is: %s", ip)
}

func main() {
	srv.Addr = port
	http.HandleFunc("/ip", getClientIp)
	srv.ListenAndServe()
}
