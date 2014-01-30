package template

import (
	"fmt"
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
	"strings"
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
	tmpl                     *template.Template
	inj                      inject.Injector
	stylesheets, javascripts []string
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
		stylesheets: []string{
			"/stylesheets/normalize.css",
		},
		javascripts: []string{},
	}
}

func (t *Tmpl) Funcs() template.FuncMap {
	return template.FuncMap{
		"noescape": func(input string) template.HTML {
			return template.HTML(input)
		},
		"stylesheets": func() template.HTML {
			imports := make([]string, len(t.stylesheets))
			for i, s := range t.stylesheets {
				imports[i] = fmt.Sprintf(
					`<link href="%s" rel="stylesheet" type="text/css" media="all">`, s)
			}
			return template.HTML(strings.Join(imports, "\n"))
		},
		"javascripts": func() template.HTML {
			imports := make([]string, len(t.javascripts))
			for i, s := range t.javascripts {
				imports[i] = fmt.Sprintf(
					`<script src="%s"></script>`, s)
			}
			return template.HTML(strings.Join(imports, "\n"))
		},
	}
}

func (t *Tmpl) AddStylesheet(path string) {
	t.stylesheets = append(t.stylesheets, path)
}

func (t *Tmpl) AddJavascript(path string) {
	t.javascripts = append(t.javascripts, path)
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
						_, err = t.tmpl.New(rel).Funcs(t.Funcs()).Parse(string(buf))
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
