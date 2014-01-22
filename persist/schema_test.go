package persist

import (
	"testing"
)

func TestSchema(t *testing.T) {
	s := NewSchema()
	table := s.Table("test_table")
	table.AddColumn("test_column", &Column{
		Type: &Integer{},
	})
	table2 := s.Table("test_table")
	if _, ok := table2.Columns["test_column"].Type.(*Integer); !ok {
		t.Fatal("Table wasn't preserved")
	}
}
