package persist

import (
	"bytes"
	"database/sql"
	"fmt"
	"strings"
)

type DialectSqlite struct {
	Dialect
}

func (d *DialectSqlite) ColumnDefSql(name string, column *Column) (
	string, error) {
	colBuf := bytes.NewBufferString(name)
	if column.Unsigned {
		colBuf.WriteString(" UNSIGNED")
	}
	switch column.Type.(type) {
	case *Bool, *SmallInt, *Integer, *BigInt:
		colBuf.WriteString(" INTEGER")
	case *Decimal, *Float, *Double:
		colBuf.WriteString(" REAL")
	case *String, *Text:
		colBuf.WriteString(" TEXT")
	default:
		return "", fmt.Errorf("Unknown column type for %s", name)
	}
	colBuf.WriteString(d.ColumnOptionSql(column))
	return colBuf.String(), nil
}

func (d *DialectSqlite) ColumnOptionSql(column *Column) string {
	options := bytes.NewBufferString("")
	if column.PrimaryKey {
		options.WriteString(" PRIMARY KEY")
	}
	if column.AutoIncrement {
		options.WriteString(" AUTOINCREMENT")
	}
	if column.Unique {
		options.WriteString(" UNIQUE")
	}
	return options.String()
}

func (d *DialectSqlite) SyncTableSchemaSql(db *sql.DB, t *Table) (
	string, error) {
	// Check if table exists
	var name string
	err := db.QueryRow(
		"SELECT name FROM sqlite_master WHERE type='table' AND name=?;",
		t.Name).Scan(&name)
	exists := true
	if err == sql.ErrNoRows {
		exists = false
	} else if err != nil {
		return "", err
	}
	buf := bytes.NewBufferString("")
	if exists {
		// Find if we're missing any columns
		rows, err := db.Query("PRAGMA table_info(" + t.Name + ");")
		if err != nil {
			return "", err
		}
		dbTableCols := map[string]bool{}
		for rows.Next() {
			var (
				name string
				ign  sql.NullString
			)
			if err := rows.Scan(&ign, &name, &ign, &ign, &ign, &ign); err != nil {
				return "", err
			}
			dbTableCols[name] = true
		}
		missingColumns := []string{}
		for cName, _ := range t.Columns {
			if !dbTableCols[cName] {
				missingColumns = append(missingColumns, cName)
			}
		}
		if len(missingColumns) > 0 {
			for _, cName := range missingColumns {
				buf.WriteString("ALTER TABLE ")
				buf.WriteString(t.Name)
				buf.WriteString(" ADD COLUMN ")
				colDef, err := d.ColumnDefSql(cName, t.Columns[cName])
				if err != nil {
					return "", err
				}
				buf.WriteString(colDef)
				buf.WriteString(";")
			}
		}
	} else {
		// Create table and columns
		buf.WriteString("CREATE TABLE ")
		buf.WriteString(t.Name)
		buf.WriteString("(")
		colDefs := []string{}
		for cName, c := range t.Columns {
			colDef, err := d.ColumnDefSql(cName, c)
			if err != nil {
				return "", err
			}
			colDefs = append(colDefs, colDef)
		}
		buf.WriteString(strings.Join(colDefs, ", "))
		buf.WriteString(");")
	}
	// Check if we're missing any indexes
	for _, index := range t.Indexes {
		indexName := fmt.Sprintf("%s_%s", t.Name, strings.Join(index, "_"))
		var name string
		err := db.QueryRow(
			"SELECT name FROM sqlite_master WHERE type='index' AND name=?",
			indexName).Scan(&name)
		if err == sql.ErrNoRows {
			buf.WriteString(fmt.Sprintf("CREATE INDEX %s ON %s(%s);",
				indexName, t.Name, strings.Join(index, ",")))
		} else if err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}
