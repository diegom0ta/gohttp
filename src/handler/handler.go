package handler

import (
	"net/http"

	"github.com/diegom0ta/gohttp/src/helper"
)

func ClientHandler(w http.ResponseWriter, req *http.Request) {
	ip := helper.ReadClientIp(req)
	w.Write([]byte("Your IP is: " + ip))
}
