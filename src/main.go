package main

import (
	"context"
	"log"
	"net/http"
)

type server struct {
	port string
}

func newServer(port string) *server {
	srv := &server{
		port: port,
	}

	return srv
}

var port = ":8090"

var srv *http.Server

func readClientIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func closeConn(srv *http.Server) {
	ctx := context.TODO()
	srv.Shutdown(ctx)
}

func getClientIP(w http.ResponseWriter, req *http.Request) {
	if ip := readClientIP(req); ip != "" {
		log.Print("Client IP is: ", ip)
	} else {
		ctx := req.Context()
		err := ctx.Err()
		log.Fatal(err)
		srv := new(http.Server)
		closeConn(srv)
	}
}

func main() {
	server := newServer(port)

	srv = new(http.Server)
	srv.Addr = server.port

	http.HandleFunc("/ip", getClientIP)
	srv.ListenAndServe()
}
