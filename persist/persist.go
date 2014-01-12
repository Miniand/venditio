package persist

import (
	"github.com/Miniand/venditio/core"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type persistInit struct {
	dbChecked bool
}

func Register(v *core.Venditio) {
	db, err := Connect()
	if err != nil {
		panic(err.Error())
	}
	v.Map(db)
	v.Map(NewRegistry())
	v.Map(&persistInit{})
	v.Use(func(r *Registry, init *persistInit) {
		if !init.dbChecked {
			init.dbChecked = true
		}
	})
}
