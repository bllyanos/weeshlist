package render

import (
	"github.com/bllyanos/weeshlist/internal/pkg/tp"
	"net/http"
)

func NotFound(templates *tp.Template, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	err := templates.ExecuteTemplate(w, "404", templates.GetData(r, nil))
	if err != nil {
		HandleError(err, w, r)
	}
}
