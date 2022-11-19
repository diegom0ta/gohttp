package main

import (
	"log"
	"net/http"
	"os"

	"github.com/diegom0ta/gohttp/src/handler"
)

var port = os.Getenv("PORT")

const route = "/"

func main() {

	if port == "" {
		port = "9000"
	}

	mux := http.NewServeMux()
	mux.HandleFunc(route, handler.ClientHandler)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
