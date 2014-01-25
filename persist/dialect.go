package persist

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Dialect interface {
	SyncTableSchemaSql(db *sql.DB, t *Table) ([]string, error)
}

func DatabaseDialect(driver string) (Dialect, error) {
	switch driver {
	case "sqlite3":
		return &DialectSqlite{}, nil
	}
	return nil, fmt.Errorf("Could not find dialect for driver %s", driver)
}
