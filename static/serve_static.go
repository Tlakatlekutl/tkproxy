package static

import "net/http"

//Initialize static server
func init() {
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
}
