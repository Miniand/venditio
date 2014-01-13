package admin

import (
	"github.com/Miniand/venditio/core"
	"github.com/codegangsta/martini-contrib/render"
	"net/http"
)

const (
	BASE_PATH = "/admin"
)

func Register(v *core.Venditio) {
	v.Map(&Admin{})
	v.Get(BASE_PATH, func(r render.Render) {
		r.HTML(http.StatusOK, "admin", nil)
	})
}

type Admin struct{}
