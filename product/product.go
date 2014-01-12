package product

import (
	"github.com/Miniand/venditio/core"
	"github.com/Miniand/venditio/persist"
)

func Register(v *core.Venditio) {
	_, err := v.Invoke(func(r *persist.Registry) {
		t := r.Table("products")
		t.AddColumn("name", &persist.String{})
		t.AddIndex("name")
	})
	if err != nil {
		panic(err.Error())
	}
}
