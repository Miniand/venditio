package template

import (
	"github.com/Miniand/venditio/asset"
	"github.com/Miniand/venditio/config"
	"github.com/Miniand/venditio/core"
	"github.com/Miniand/venditio/inject"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

const (
	DEP_TEMPLATE           = "template"
	ASSET_BASE_DIR         = "templates"
	CONFIG_CACHE_TEMPLATES = "CACHE_TEMPLATES"
)

type Templater interface {
	Render(wr io.Writer, tmpl string, data interface{}) error
}

type Tmpl struct {
	tmpl *template.Template
	inj  inject.Injector
}

var funcs = template.FuncMap{
	"noescape": func(input string) template.HTML {
		return template.HTML(input)
	},
}

func Register(v *core.Venditio) {
	v.MustGet(asset.DEP_ASSET).(asset.Resolver).AddPackagePath(
		"github.com/Miniand/venditio/template/assets")
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
	if t.tmpl == nil || t.inj.MustGet(config.DEP_CONFIG).(*config.Config).Get(
		CONFIG_CACHE_TEMPLATES) != "1" {
		t.Compile()
	}
	return t.tmpl.ExecuteTemplate(wr, tmpl, data)
}

func (t *Tmpl) Compile() {
	t.tmpl = template.New("venditio")
	for _, rp := range t.inj.MustGet(
		asset.DEP_ASSET).(asset.Resolver).ResolvePaths() {
		dir := path.Join(rp, ASSET_BASE_DIR)
		if _, err := os.Stat(dir); !os.IsNotExist(err) {
			filepath.Walk(dir, func(path string, info os.FileInfo,
				err error) error {
				if filepath.Ext(path) == ".tmpl" {
					rel, err := filepath.Rel(dir, path)
					if err != nil {
						return err
					}
					if t.tmpl.Lookup(rel) == nil {
						buf, err := ioutil.ReadFile(path)
						if err != nil {
							return err
						}
						_, err = t.tmpl.New(rel).Funcs(funcs).Parse(string(buf))
						if err != nil {
							return err
						}
					}
				}
				return nil
			})
		}
	}
}
