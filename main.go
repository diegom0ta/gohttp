package main

import (
	"fmt"
	"net/http"
)

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
	ip := readUserIP(req)
	fmt.Println(ip)
}

func main() {
	http.HandleFunc("/myip", getClientIP)
	http.ListenAndServe(":8090", nil)
}
