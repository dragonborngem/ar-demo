package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func route() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.PathPrefix("/demo/").Handler(http.StripPrefix("/demo/", http.FileServer(http.Dir("./"))))
	return r
}

func main() {

	r := route()
	http.Handle("/", r)
	if os.Getenv("MODE") != "" {
		log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
	} else {
		log.Fatal(http.ListenAndServe(":"+"10000", nil))
	}
}
