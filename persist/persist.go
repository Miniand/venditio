package persist

import (
	"fmt"
	"github.com/Miniand/venditio/core"
	"github.com/Miniand/venditio/inject"
)

const (
	DEP_DB     = "persistDb"
	DEP_SCHEMA = "persistSchema"
)

type persistInit struct {
	dbChecked bool
}

func Register(v *core.Venditio) {
	v.BindFactory(DEP_SCHEMA, func(i inject.Injector) interface{} {
		return NewSchema()
	})
	v.BindFactory(DEP_DB, func(i inject.Injector) interface{} {
		// Connect
		db, err := Connect()
		if err != nil {
			panic(err.Error())
		}
		// Check schema
		s := i.MustGet(DEP_SCHEMA).(*Schema)
		schemaSql, err := SyncSchemaSql(db, s)
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
		return db
	})
}
