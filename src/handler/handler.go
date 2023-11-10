package handler

import (
	"fmt"
	"net/http"

	"github.com/diegom0ta/gohttp/src/helper"
)

func ClientHandler(w http.ResponseWriter, req *http.Request) {
	ip := helper.ReadClientIp(req)
	w.Write([]byte(fmt.Sprintf("Your IP is: %s\n", ip)))
}
