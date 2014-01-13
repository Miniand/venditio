package view

import (
	"github.com/Miniand/venditio/core"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"html/template"
)

func Register(v *core.Venditio) {
	v.Use(render.Renderer(render.Options{
		Funcs: []template.FuncMap{
			template.FuncMap{
				"noescape": func(input string) template.HTML {
					return template.HTML(input)
				},
			},
		},
	}))
	v.Use(martini.Static("assets", martini.StaticOptions{
		Prefix: "assets/",
	}))
}
