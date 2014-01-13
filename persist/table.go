package persist

import (
	"database/sql"
)

type Table struct {
	Name    string
	Columns map[string]*Column
	Indexes [][]string
}

func NewTable(name string) *Table {
	return &Table{
		Name: name,
		Columns: map[string]*Column{
			"id": &Column{
				Type:          &Integer{},
				PrimaryKey:    true,
				AutoIncrement: true,
			},
		},
		Indexes: [][]string{},
	}
}

func (t *Table) AddColumn(name string, column *Column) {
	t.Columns[name] = column
}

func (t *Table) AddIndex(index []string) {
	t.Indexes = append(t.Indexes, index)
}

func (t *Table) CheckDb(db *sql.DB) (string, error) {
	return "", nil
}
