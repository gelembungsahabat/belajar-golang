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
	ID    string `json:"blog_id"`
	Judul string `json:"title"`
	Isi   string `json:"isi"`
}

//fungsi utama
func main() {
	blogs = []blog{
		blog{ID: "1", Judul: "haha", Isi: "haha"},
		blog{ID: "2", Judul: "haha", Isi: "haha"},
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
		r.Get("/{ID}", singlePostingan)
		r.Post("/", createPostingan)
		r.Delete("/{ID}", deletePostingan)
		r.Put("/{ID}", updatePostingan)

	})

	http.ListenAndServe(":1320", r)
}

//fungsi untuk menampilkan homepage(root)
func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "halo ahaha")
	fmt.Println("hahaha")
}

//fungsi untuk menampilkan semua data dari array 'blogs'
func allPostingan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("endpoint hit: returnAllPostingan")
	json.NewEncoder(w).Encode(blogs)
}

//fungsi untuk menampilkan data array 'blogs' berdasarkan ID
func singlePostingan(w http.ResponseWriter, r *http.Request) {
	fmt.Println("endpoint hit: returnSinglePostingan")

	w.Header().Set("Content-Type", "application/json")
	for _, postingan := range blogs {
		if postingan.ID == chi.URLParam(r, "ID") {
			json.NewEncoder(w).Encode(postingan)
		}
	}
}

//fungsi untuk menambahkan data ke array 'blogs'
func createPostingan(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	var Blog blog
	json.Unmarshal(data, &Blog)

	if Blog.ID = strconv.Itoa(len(blogs)); Blog.ID != "0" {
		index, err := strconv.Atoi(blogs[len(blogs)-1].ID)
		if err != nil {
			panic(err)
		}
		Blog.ID = strconv.Itoa(index + 1)
	}

	blogs = append(blogs, Blog)
	json.NewEncoder(w).Encode(Blog.ID)

}

//fungsi untuk DELETE
func deletePostingan(w http.ResponseWriter, r *http.Request) {
	for index, blog := range blogs {
		if blog.ID == chi.URLParam(r, "ID") {
			blogs = append(blogs[:index], blogs[index+1:]...)
			fmt.Printf("DELETE DATA")
			json.NewEncoder(w).Encode("Berhasil hore")
		}
	}
}

func updatePostingan(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	var Blog blog
	json.Unmarshal(data, &Blog)
	bid := chi.URLParam(r, "ID")

	for index, b := range blogs {
		if b.ID == bid {
			blogs[index] = blog{
				ID:    b.ID,
				Judul: Blog.Judul,
				Isi:   Blog.Isi,
			}
		}
	}
}
