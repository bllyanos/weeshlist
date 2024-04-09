package render

import (
	"net/http"

	"github.com/bllyanos/weeshlist/internal/pkg/tp"
)

func Page(templates *tp.Template, w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	err := templates.ExecuteTemplate(w, name, templates.GetData(r, data))
	if err != nil {
		HandleError(err, w, r)
	}
}
