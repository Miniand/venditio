package persist

import (
	"database/sql"
	"github.com/Miniand/venditio/model"
	"os"
	"strings"
)

func DatabaseDriver() string {
	driver := os.Getenv("DATABASE_DRIVER")
	if driver == "" {
		driver = "sqlite3"
	}
	return driver
}

func DatabaseOptions() string {
	options := os.Getenv("DATABASE_OPTIONS")
	if options == "" {
		options = "database.sqlite"
	}
	return options
}

func Connect() (*sql.DB, error) {
	return sql.Open(DatabaseDriver(), DatabaseOptions())
}

func SyncRegistrySchemaSql(db *sql.DB, r *Registry) (string, error) {
	dialect, err := DatabaseDialect(DatabaseDriver())
	if err != nil {
		return "", err
	}
	tSql := []string{}
	for _, t := range r.Tables {
		s, err := dialect.SyncTableSchemaSql(db, t)
		if err != nil {
			return "", err
		}
		tSql = append(tSql, s)
	}
	return strings.Join(tSql, ""), nil
}

func RowsToModels(rows *sql.Rows, t *Table) (models []model.Model, err error) {
	var cols []string
	for rows.Next() {
		if cols == nil {
			if cols, err = rows.Columns(); err != nil {
				return
			}
		}
		m := model.Model{}
		scans := make([]interface{}, len(cols))
		for i, c := range cols {
			if t.Columns[c] != nil {
				scans[i] = t.Columns[c].Type.RawType()
			} else {
				scans[i] = &sql.NullString{}
			}
		}
		if err = rows.Scan(scans...); err != nil {
			return
		}
		for i, c := range cols {
			switch dt := scans[i].(type) {
			case *sql.NullBool:
				if dt.Valid {
					m[c] = dt.Bool
				}
			case *sql.NullInt64:
				if dt.Valid {
					m[c] = dt.Int64
				}
			case *sql.NullFloat64:
				if dt.Valid {
					m[c] = dt.Float64
				}
			case *sql.NullString:
				if dt.Valid {
					m[c] = dt.String
				}
			default:
				if scans[i] != nil {
					m[c] = scans[i]
				}
			}
		}
		models = append(models, m)
	}
	return
}
