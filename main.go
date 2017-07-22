package main

import (
	"net/http"

	_ "github.com/Tlakatlekutl/tkproxy/static"
	"github.com/Tlakatlekutl/tkproxy/log"
	_ "github.com/Tlakatlekutl/tkproxy/server"
)

func main() {
	log.SetDebugLevel("ALL")

	locationUrl := "localhost:7593"
	log.Info("Start server on %s", locationUrl)

	http.ListenAndServe(locationUrl, nil)
}
