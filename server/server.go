package server

import (
	"net/http"
	"html/template"
	"encoding/json"
	"fmt"

	"github.com/Tlakatlekutl/tkproxy/log"
	"github.com/Tlakatlekutl/tkproxy/parser"
	"io/ioutil"
	"net/url"
)

func init() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/go/", OpenSiteHandler)
	http.HandleFunc("/source", ResourceHandler)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	log.Trace("%v  %+v", r.Method, r.URL)

	t, err := template.ParseFiles("server/templates/index.tmpl")
	if err != nil {
		log.Error(err.Error())
		return
	}

	if err = t.Execute(w, template.HTML("<h1>Welcome </h1>")); err != nil {
		log.Error(err.Error())
		return
	}
}

func OpenSiteHandler(w http.ResponseWriter, r *http.Request)  {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var u struct{
		Url string `json:"url"`
	}
	err := decoder.Decode(&u)
	if err != nil {
		panic(err)
	}

	log.Trace("%v %+v,  goto: %+v", r.Method, r.URL, u.Url)

	resp, err := http.Get(u.Url)
	if err != nil {
		fmt.Fprintf(w, "Get %s returned error %s", u.Url, err.Error() )
		log.Error(err.Error())
		return
	}
	urlParsed, _ := url.Parse(u.Url)

	body, err := parser.ChangeSourceUrl(resp.Body, "/source?from=http://", urlParsed.Host)
	defer resp.Body.Close()

	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
	w.Write([]byte(body))
}

func ResourceHandler(w http.ResponseWriter, r *http.Request)  {
	log.Trace("%v  %+v", r.Method, r.URL)
	source, ok := r.URL.Query()["from"]
	if !ok {
		log.Error("Error parse \"from\" get parametr from utl")
		http.NotFound(w, r)
	}
	resp, err := http.Get(source[0])
	if err != nil {
		fmt.Fprintf(w, "Get %s returned error %s", source[0], err.Error() )
		log.Error(err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	for name, values := range resp.Header {
		w.Header()[name] = values
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}