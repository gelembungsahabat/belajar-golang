package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

//'blogs' sebagai array dari 'blog'
var blogs []blog

//struktur data dari variabel 'blog'
type blog struct {
	Id    string `"json":"ID"`
	Judul string `"json":"Title"`
	Isi   string `"json":"Isi"`
}

//fungsi utama
func main() {
	blogs = []blog{
		blog{Id: "1", Judul: "haha", Isi: "haha"},
		blog{Id: "2", Judul: "haha", Isi: "haha"},
	}
	handleRequest()
}

//fungsi untuk routing
func handleRequest() {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/", homePage)

	r.Route("/blog", func(r chi.Router) {
		r.Get("/", allPostingan)
		r.Get("/{id}", singlePostingan)
		r.Post("/", createPostingan)
	})

	http.ListenAndServe(":1320", r)
}

//fungsi untuk menampilkan homepage(root)
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "halo ahaha")
	fmt.Println("hahaha")
}

//fungsi untuk menampilkan semua data dari array 'blogs'
func allPostingan(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint hit: returnAllPostingan")
	json.NewEncoder(w).Encode(blogs)
}

//fungsi untuk menampilkan data array 'blogs' berdasarkan ID
func singlePostingan(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint hit: returnSinglePostingan")

	for _, postingan := range blogs {
		if postingan.Id == chi.URLParam(r, "id") {
			json.NewEncoder(w).Encode(postingan)
		}
	}
}

//fungsi untuk menambahkan data ke array 'blogs'
func createPostingan(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	var Blog blog
	json.Unmarshal(data, &Blog)

	if Blog.Id = strconv.Itoa(len(blogs)); Blog.Id != "0" {
		index, err := strconv.Atoi(blogs[len(blogs)-1].Id)
		if err != nil {
			panic(err)
		}
		Blog.Id = strconv.Itoa(index + 1)
	}

	blogs = append(blogs, Blog)
	json.NewEncoder(w).Encode(Blog.Id)

}
