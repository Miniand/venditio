package persist

import (
	"database/sql"
)

type Table struct {
	Name          string
	PrimaryKey    string
	Columns       map[string]Typeable
	AutoIncrement map[string]bool
	Index         map[string]bool
	Unique        map[string]bool
	NotNull       map[string]bool
}

func NewTable(name string) *Table {
	return &Table{
		Name:       name,
		PrimaryKey: "id",
		Columns: map[string]Typeable{
			"id": &Integer{},
		},
		Index: map[string]bool{
			"id": true,
		},
	}
}

func (t *Table) AddIndex(column string) {
	t.Index[column] = true
}

func (t *Table) RemoveIndex(column string) {
	t.Index[column] = false
}

func (t *Table) AddUnique(column string) {
	t.Unique[column] = true
}

func (t *Table) RemoveUnique(column string) {
	t.Unique[column] = false
}

func (t *Table) AddAutoIncrement(column string) {
	t.AutoIncrement[column] = true
}

func (t *Table) RemoveAutoIncrement(column string) {
	t.AutoIncrement[column] = false
}

func (t *Table) AddNotNull(column string) {
	t.NotNull[column] = true
}

func (t *Table) RemoveNotNull(column string) {
	t.NotNull[column] = false
}

func (t *Table) AddColumn(name string, kind Typeable) {
	t.Columns[name] = kind
}

func (t *Table) CheckDb(db *sql.DB) (string, error) {
	return "", nil
}
