package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/html/index.html")
}

func getJS(w http.ResponseWriter, r *http.Request) {
	filename := strings.Split(r.URL.Path, "/")[2]
	http.ServeFile(w, r, "web/js/"+filename)
}

func getCSS(w http.ResponseWriter, r *http.Request) {
	filename := strings.Split(r.URL.Path, "/")[2]
	http.ServeFile(w, r, "web/css/"+filename)
}

func getImg(w http.ResponseWriter, r *http.Request) {
	filename := strings.Split(r.URL.Path, "/")[2]
	http.ServeFile(w, r, "web/img/"+filename)
}

func RunWebserver() {
	// Setup and run all routes
	fmt.Println("Starting web server")

	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/js/{filename}", getJS)
	r.HandleFunc("/css/{filename}", getCSS)
	r.HandleFunc("/img/{filename}", getImg)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
