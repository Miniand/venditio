package venditio

import (
	"github.com/Miniand/venditio/core"
	"github.com/Miniand/venditio/persist"
	"github.com/Miniand/venditio/product"
)

func New() *core.Venditio {
	v := core.New()
	persist.Register(v)
	product.Register(v)
	return v
}
