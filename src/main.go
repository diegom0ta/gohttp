package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Server listening on port: %s\n", port)

	go func() {
		if err := srv.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Listen and serve error: %v\n", err)
		}
	}()

	<-ctx.Done()

	log.Println("Got interruption signal")

	if err := srv.Shutdown(context.TODO()); err != nil {
		log.Printf("Server shutdown returned an error: %v\n", err)
	}
}
