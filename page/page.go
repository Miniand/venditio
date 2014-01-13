package page

import (
	"database/sql"
	"github.com/Miniand/venditio/core"
	"github.com/Miniand/venditio/persist"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"net/http"
)

func Register(v *core.Venditio) {
	v.Invoke(func(reg *persist.Registry) {
		t := reg.Table("pages")
		t.AddColumn("url", &persist.Column{
			Type:    &persist.String{},
			NotNull: true,
		})
		t.AddColumn("title", &persist.Column{
			Type:    &persist.String{},
			NotNull: true,
		})
		t.AddColumn("body", &persist.Column{
			Type:    &persist.Text{},
			NotNull: true,
		})
		t.AddIndex([]string{"title"})
	})
	v.Get("/page/:url", func(db *sql.DB, reg *persist.Registry,
		r render.Render, params martini.Params) {
		rows, err := db.Query("SELECT * FROM pages WHERE url=?",
			params["url"])
		if err != nil {
			panic(err.Error())
		}
		models, err := persist.RowsToModels(rows, reg.Table("pages"))
		if err != nil {
			panic(err.Error())
		}
		if len(models) == 0 {
			r.Redirect("/")
		}
		r.HTML(http.StatusOK, "page", models[0])
	})
}
