package weeshlist

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bllyanos/weeshlist/pkg/multiplexer"
)

func applyLoggerMiddleware(rootMux *multiplexer.Mulx) {
	rootMux.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()

			// execute http
			next.ServeHTTP(w, r)

			method := r.Method
			path := r.URL.Path
			timeFormatted := startTime.Format(time.RFC3339)

			fmt.Printf("%s %s %s -- (%s)\n", timeFormatted, method, path, time.Since(startTime))
			return
		})
	})
}

func applyStaticFileServer(rootMux *multiplexer.Mulx) {
	rootMux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
}
