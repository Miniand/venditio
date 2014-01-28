package template

import (
	"errors"
	"github.com/Miniand/venditio/asset"
	"github.com/Miniand/venditio/core"
	"github.com/Miniand/venditio/inject"
	"html/template"
	"io"
	"io/ioutil"
	"path"
)

const (
	DEP_TEMPLATE   = "template"
	ASSET_BASE_DIR = "templates"
)

type Templater interface {
	Render(wr io.Writer, tmpl string, data interface{}) error
}

type Tmpl struct {
	inj inject.Injector
}

func Register(v *core.Venditio) {
	v.BindFactory(DEP_TEMPLATE, func(i inject.Injector) interface{} {
		return New(i)
	})
}

func New(inj inject.Injector) Templater {
	return &Tmpl{
		inj: inj,
	}
}

func (t *Tmpl) Render(wr io.Writer, tmpl string, data interface{}) error {
	f := t.inj.MustGet(asset.DEP_ASSET).(asset.Resolver).Resolve(
		path.Join(ASSET_BASE_DIR, tmpl))
	if len(f) == 0 {
		return errors.New("Could not find template")
	}
	raw, err := ioutil.ReadFile(f[0])
	parser := template.New(tmpl)
	parser.Funcs(template.FuncMap{
		"noescape": func(input string) template.HTML {
			return template.HTML(input)
		},
	})
	_, err = parser.Parse(string(raw))
	if err != nil {
		return err
	}
	return parser.Execute(wr, data)
}
