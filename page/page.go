package page

import (
	"database/sql"
	"github.com/Miniand/venditio/core"
	"github.com/Miniand/venditio/persist"
	"github.com/Miniand/venditio/web"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func Register(v *core.Venditio) {
	// Schema
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
	// Web
	router := v.MustGet(web.DEP_ROUTER).(*mux.Router)
	router.HandleFunc("/pages/{url}", func(w http.ResponseWriter,
		r *http.Request) {
		db := v.MustGet(persist.DEP_DB).(*sql.DB)
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
		io.WriteString(w, models[0]["body"].(string))
	})
}
