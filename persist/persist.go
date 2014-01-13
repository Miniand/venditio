package persist

import (
	"database/sql"
	"fmt"
	"github.com/Miniand/venditio/core"
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
	v.Use(func(r *Registry, init *persistInit, db *sql.DB) {
		if !init.dbChecked {
			schemaSql, err := SyncRegistrySchemaSql(db, r)
			if err != nil {
				panic(err.Error())
			}
			if schemaSql != "" {
				fmt.Printf("Running following schema update on database:\n%s\n",
					schemaSql)
				if _, err := db.Exec(schemaSql); err != nil {
					panic(err.Error())
				}
			}
			init.dbChecked = true
		}
	})
}
