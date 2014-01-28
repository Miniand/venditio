package page

import (
	"database/sql"
	"github.com/Miniand/venditio/asset"
	"github.com/Miniand/venditio/core"
	"github.com/Miniand/venditio/persist"
	"github.com/Miniand/venditio/template"
	"github.com/Miniand/venditio/web"
	"github.com/gorilla/mux"
	"net/http"
)

func Register(v *core.Venditio) {
	registerSchema(v)
	registerAssets(v)
	registerRoutes(v)
}

func registerSchema(v *core.Venditio) {
	schema := v.MustGet(persist.DEP_SCHEMA).(*persist.Schema)
	t := schema.Table("pages")

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
}

func registerAssets(v *core.Venditio) {
	v.MustGet(asset.DEP_ASSET).(asset.Resolver).AddPackagePath(
		"github.com/Miniand/venditio/page/assets")
}

func registerRoutes(v *core.Venditio) {
	schema := v.MustGet(persist.DEP_SCHEMA).(*persist.Schema)
	router := v.MustGet(web.DEP_ROUTER).(*mux.Router)
	router.HandleFunc("/pages/{url}", func(w http.ResponseWriter,
		r *http.Request) {
		db := v.MustGet(persist.DEP_DB).(*sql.DB)
		tmpl := v.MustGet(template.DEP_TEMPLATE).(template.Templater)
		vars := mux.Vars(r)
		rows, err := db.Query("SELECT * FROM pages WHERE url=?",
			vars["url"])
		if err != nil {
			panic(err.Error())
		}
		models, err := persist.RowsToModels(rows, schema.Table("pages"))
		if err != nil {
			panic(err.Error())
		}
		if len(models) == 0 {
			return
		}
		err = tmpl.Render(w, "page.tmpl", map[string]interface{}{
			"title": models[0]["title"],
			"body":  models[0]["body"],
		})
		if err != nil {
			panic(err.Error())
		}
	})
}
