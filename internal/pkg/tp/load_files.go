package tp

import (
	"fmt"
	"html/template"
	"time"

	"github.com/bllyanos/weeshlist/internal/pkg/comm"
)

func getTemplateFuncs() template.FuncMap {
	return template.FuncMap{
		"FormatTime": func(t time.Time) string {
			return t.Format(time.RFC3339)
		},
	}
}

func LoadTemplates(path string, defaultMeta comm.JSON) *Template {
	tpl, err := template.New("base").Funcs(getTemplateFuncs()).ParseGlob("web/template/**/*.html")
	if err != nil {
		fmt.Println("error loading templates:", err)
		panic(err)
	}
	wrap := NewTemplate(tpl, defaultMeta)
	return &wrap
}
