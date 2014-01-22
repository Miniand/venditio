package persist

type Schema struct {
	Tables map[string]*Table
}

func NewSchema() *Schema {
	return &Schema{
		Tables: map[string]*Table{},
	}
}

func (s *Schema) Table(name string) *Table {
	t, ok := s.Tables[name]
	if !ok {
		t = NewTable(name)
		s.Tables[name] = t
	}
	return t
}
