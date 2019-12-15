//
// REST
// ====
// This example demonstrates a HTTP REST web service with some fixture data.
// Follow along the example and patterns.
//
// Boot the server:
// ----------------
// $ go run *.go
//
// Client requests:
// ----------------
// $ curl http://localhost:3333/
// Halo Dunia!.
//
// $ curl http://localhost:3333/data
// [{"id":"1","title":"Hi"},{"id":"2","title":"sup"}]
//
// $ curl http://localhost:3333/data/1
// {"id":"1","title":"Hi"}
//
// $ curl -X DELETE http://localhost:3333/data/1
// {"id":"1","title":"Hi"}
//
// $ curl http://localhost:3333/data/1
// "Not Found"
//
// $ curl -X POST -d '{"id":"will-be-omitted","title":"awesomeness"}' http://localhost:3333/data
// {"id":"97","title":"awesomeness"}
//
// $ curl http://localhost:3333/data/97
// {"id":"97","title":"awesomeness"}
//
// $ curl http://localhost:3333/data
// [{"id":"2","title":"sup"},{"id":"97","title":"awesomeness"}]
//
package main

import (
	"flag"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

var routes = flag.Bool("routes", false, "Generate router documentation")

func main() {
	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Halo Dunia!."))
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})

	// RESTy routes for "data" resource
	r.Route("/data", func(r chi.Router) {
		r.With(paginate).Get("/", DataList)
		r.Post("/", CreateArticle)       // POST /data
		r.Get("/search", SearchArticles) // GET /data/search

		r.Route("/{data_ID}", func(r chi.Router) {
			r.Use(ArticleCtx)            // Load the *Article on the request context
			r.Get("/", GetArticle)       // GET /data/123
			r.Put("/", UpdateArticle)    // PUT /data/123
			r.Delete("/", DeleteArticle) // DELETE /data/123
		})

		// GET /data/whats-up
		// r.With(ArticleCtx).Get("/{dataSlug:[a-z-]+}", GetArticle)
	})

	// Mount the admin sub-router, which btw is the same as:
	// r.Route("/admin", func(r chi.Router) { admin routes here })
	r.Mount("/admin", adminRouter())

	http.ListenAndServe(":3333", r)
}
