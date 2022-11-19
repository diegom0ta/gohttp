package helper

import "net/http"

func ReadClientIp(r *http.Request) string {
	ipAddr := r.Header.Get("X-Real-Ip")
	if ipAddr == "" {
		ipAddr = r.Header.Get("X-Forwarded-For")
	}
	if ipAddr == "" {
		ipAddr = r.RemoteAddr
	}
	return ipAddr
}
