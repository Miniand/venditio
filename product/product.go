package product

import (
	"database/sql"
	"fmt"
	"github.com/Miniand/venditio/core"
	"github.com/Miniand/venditio/persist"
)

func Register(v *core.Venditio) {
	_, err := v.Invoke(func(r *persist.Registry) {
		t := r.Table("products")
		t.AddColumn("name", &persist.Column{
			Type: &persist.String{},
		})
		t.AddColumn("price", &persist.Column{
			Type: &persist.Float{},
		})
		t.AddColumn("enabled", &persist.Column{
			Type: &persist.Bool{},
		})
		t.AddIndex([]string{"name"})
	})
	if err != nil {
		panic(err.Error())
	}
	v.Use(func(db *sql.DB, r *persist.Registry) {
		rows, err := db.Query("SELECT * FROM products;")
		if err != nil {
			panic(err.Error())
		}
		models, err := persist.RowsToModel(rows, r.Table("products"))
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%#v\n", models[0]["name"])
		fmt.Printf("%#v\n", models[0]["enabled"])
		fmt.Printf("%#v\n", models[0]["price"])
	})
}
