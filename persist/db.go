package persist

import (
	"database/sql"
	"os"
)

const (
	DB_TABLE_MODIFICATION_TYPE_CREATE = iota
	DB_TABLE_MODIFICATION_TYPE_ALTER
)

type DbTableModification struct {
	Table   *Table
	Type    int
	Columns []string
}

func Connect() (*sql.DB, error) {
	driver := os.Getenv("DATABASE_DRIVER")
	if driver == "" {
		driver = "sqlite3"
	}
	options := os.Getenv("DATABASE_OPTIONS")
	if options == "" {
		options = "database.sqlite"
	}
	return sql.Open(driver, options)
}
