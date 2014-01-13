package product

import (
	"github.com/Miniand/venditio/core"
	"github.com/Miniand/venditio/persist"
)

func Register(v *core.Venditio) {
	_, err := v.Invoke(func(r *persist.Registry) {
		t := r.Table("products")
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
	})
	if err != nil {
		panic(err.Error())
	}
}
