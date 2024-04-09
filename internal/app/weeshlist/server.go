package weeshlist

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bllyanos/weeshlist/internal/pkg/comm"
	"github.com/bllyanos/weeshlist/internal/pkg/comm/render"
	"github.com/bllyanos/weeshlist/internal/pkg/tp"
	"github.com/bllyanos/weeshlist/pkg/multiplexer"
)

func StartServer() {

	templates := tp.LoadTemplates("web/template/**/*.html", comm.JSON{
		"Version": "0.0.1",
	})

	mux := multiplexer.NewMulx()

	// route - dash
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			render.NotFound(templates, w, r)
			return
		}
		data := comm.JSON{
			"Wishes": []comm.JSON{
				{"Name": "Wish 1", "Description": "lorem ipsum bla bla bla bla bla", "Price": 100_000, "URL": "https://google.com/"},
				{"Name": "Wish 2", "Description": "lorem ipsum bla bla bla bla bla", "Price": 200_000},
				{"Name": "Wish 3", "Description": "lorem ipsum bla bla bla bla bla", "Price": 300_000},
				{"Name": "Wish 4", "Description": "lorem ipsum bla bla bla bla bla", "Price": 400_000},
				{"Name": "Wish 4", "Description": "lorem ipsum bla bla bla bla bla", "Price": 400_000},
				{"Name": "Wish 4", "Description": "lorem ipsum bla bla bla bla bla", "Price": 400_000},
				{"Name": "Wish 4", "Description": "lorem ipsum bla bla bla bla bla", "Price": 400_000},
			},
		}
		render.Page(templates, w, r, "index", data)
	})

	// serve static
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// catch 404
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render.NotFound(templates, w, r)
	})

	// logger middleware
	mux.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			next.ServeHTTP(w, r)
			method := r.Method
			path := r.URL.Path
			fmt.Printf("request %s %s -- took %s\n", method, path, time.Since(startTime))
			return
		})
	})

	fmt.Println("starting server on :8080")
	http.ListenAndServe("0.0.0.0:8080", mux)
}
