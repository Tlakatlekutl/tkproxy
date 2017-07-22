package server

import (
	"net/http"
	"github.com/Tlakatlekutl/tkproxy/log"
)

func responseWithError(msg string, w http.ResponseWriter, r *http.Request) {
	log.Error(msg)
	w.Write([]byte(msg))

}