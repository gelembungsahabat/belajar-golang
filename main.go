package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type blog struct {
	Id    string `"json":"ID"`
	Judul string `"json":"Title"`
	Isi   string `"json":"Isi"`
}

var Blog []blog

func main() {
	Blog = []blog{
		blog{Id: "1", Judul: "haha", Isi: "haha"},
		blog{Id: "2", Judul: "haha", Isi: "haha"},
	}
	handleRequest()
}

func handleRequest() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", homePage)
	r.HandleFunc("/blog", returnAllPostingan)
	r.HandleFunc("/blog/{id}", returnSinglePostingan)

	log.Fatal(http.ListenAndServe(":1320", r))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "halo ahaha")
	fmt.Println("hahaha")
}

func returnAllPostingan(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint hit: returnAllPostingan")
	json.NewEncoder(w).Encode(Blog)
}

func returnSinglePostingan(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint hit: returnSinglePostingan")
	variabel := mux.Vars(r)
	key := variabel["id"]

	for _, postingan := range Blog {
		if postingan.Id == key {
			json.NewEncoder(w).Encode(postingan)
		}
	}
}
