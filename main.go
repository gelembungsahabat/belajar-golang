package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func main() {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	//GET
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Wildan Ganteng!"))
	})

	r.Get("/Open", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Selamat Datang di Blog Kami !"))
	})

	//GET with ID
	r.Route("/Blog", func(r chi.Router) {
		r.With(paginate).Get("/", DataList)
		r.Post("/", CreateData)

		r.Route("/{data_ID}", func(r chi.Router) {

			r.Use(DataContext)  // Load the *Article on the request context
			r.Get("/", GetData) // GET /Blog/123
		})
	})

	http.ListenAndServe(":1320", r)
}
