package venditio

import (
	// "github.com/Miniand/venditio/admin"
	"github.com/Miniand/venditio/asset"
	"github.com/Miniand/venditio/cmd"
	"github.com/Miniand/venditio/config"
	"github.com/Miniand/venditio/core"
	"github.com/Miniand/venditio/page"
	"github.com/Miniand/venditio/persist"
	"github.com/Miniand/venditio/product"
	"github.com/Miniand/venditio/template"
	"github.com/Miniand/venditio/web"
)

func New() *core.Venditio {
	v := core.New()
	config.Register(v)
	cmd.Register(v)
	// admin.Register(v)
	web.Register(v)
	asset.Register(v)
	template.Register(v)
	persist.Register(v)
	product.Register(v)
	page.Register(v)
	return v
}
