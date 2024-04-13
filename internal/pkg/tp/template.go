package tp

import (
	"html/template"
	"net/http"
	"time"

	"github.com/bllyanos/weeshlist/internal/pkg/comm"
	"github.com/bllyanos/weeshlist/internal/pkg/comm/util"
)

type TemplateData struct {
	Meta interface{}
	Data interface{}
}

type Template struct {
	template.Template
	defaultMeta comm.JSON
}

func NewTemplate(template *template.Template, defaultMeta comm.JSON) Template {
	return Template{
		Template:    *template,
		defaultMeta: defaultMeta,
	}
}

func (t *Template) GetData(r *http.Request, meta interface{}, data interface{}) TemplateData {

	metaClone := util.CopyMap(t.defaultMeta)
	metaClone = assignTime(metaClone)
	metaClone = assignPath(metaClone, r.URL.Path)

	if meta != nil {
		util.PatchMap(metaClone, meta.(comm.JSON))
	}

	return TemplateData{
		Meta: metaClone,
		Data: data,
	}
}

func assignTime(meta comm.JSON) comm.JSON {
	meta["Time"] = time.Now()
	return meta
}

func assignPath(meta comm.JSON, path string) comm.JSON {
	meta["Path"] = path
	return meta
}
