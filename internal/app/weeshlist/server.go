package weeshlist

import (
	"fmt"
	"net/http"

	"github.com/bllyanos/weeshlist/internal/app/weeshlist/handlers"
	"github.com/bllyanos/weeshlist/internal/pkg/comm"
	"github.com/bllyanos/weeshlist/internal/pkg/tp"
	"github.com/bllyanos/weeshlist/pkg/multiplexer"
)

func StartServer() {
	// inititate global dependencies
	templates := tp.LoadTemplates("web/template/**/*.html", comm.JSON{
		"Version": "0.0.1",
	})

	// root multiplexer
	rootMux := multiplexer.NewMulx()

	// apply global middlewares and static file server
	applyLoggerMiddleware(rootMux)
	applyStaticFileServer(rootMux)

	// route - index
	indexHandler := handlers.NewIndexHandler(templates)
	indexMux := multiplexer.NewMulx()
	indexMux.HandleFunc("GET /", indexHandler.IndexHandler)
	// / not found
	indexMux.HandleFunc("/", indexHandler.NotFoundHandler)
	rootMux.Handle("/", indexMux)

	// route - auth
	authHandler := handlers.NewAuthHandler(templates)
	authMux := multiplexer.NewMulx()
	authMux.HandleFunc("GET /auth/login", authHandler.LoginPage)
	authMux.HandleFunc("POST /auth/login/submit", authHandler.LoginAction)
	authMux.HandleFunc("GET /auth/register", authHandler.RegisterPage)
	authMux.HandleFunc("POST /auth/register", authHandler.RegisterAction)
	// /auth/ not found
	authMux.HandleFunc("/auth/", indexHandler.NotFoundHandler)
	rootMux.Handle("/auth/", authMux)

	// route - search
	searchHandler := handlers.NewSearchHandler(templates)
	searchMux := multiplexer.NewMulx()
	searchMux.HandleFunc("GET /search", searchHandler.SearchHandler)
	// /search not found
	searchMux.HandleFunc("/search", indexHandler.NotFoundHandler)
	rootMux.Handle("/search", searchMux)

	fmt.Println("starting server on :8080")
	http.ListenAndServe("0.0.0.0:8080", rootMux)
}
