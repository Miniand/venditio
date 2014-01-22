package product

import (
	"github.com/Miniand/venditio/core"
	"github.com/Miniand/venditio/persist"
	"github.com/Miniand/venditio/web"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func Register(v *core.Venditio) {
	schema := v.MustGet(persist.DEP_SCHEMA).(*persist.Schema)
	t := schema.Table("products")
	t.AddColumn("name", &persist.Column{
		Type:    &persist.String{},
		NotNull: true,
	})
	t.AddColumn("price", &persist.Column{
		Type:    &persist.Float{},
		NotNull: true,
	})
	t.AddColumn("enabled", &persist.Column{
		Type:    &persist.SmallInt{},
		NotNull: true,
	})
	t.AddIndex([]string{"name"})
	router := v.MustGet(web.DEP_ROUTER).(*mux.Router)
	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Products!")
	})
}
