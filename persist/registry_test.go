package persist

import (
	"testing"
)

func TestRegistry(t *testing.T) {
	r := NewRegistry()
	table := r.Table("test_table")
	table.AddColumn("test_column", &Integer{})
	table2 := r.Table("test_table")
	if _, ok := table2.Columns["test_column"].(*Integer); !ok {
		t.Fatal("Table wasn't preserved")
	}
}
