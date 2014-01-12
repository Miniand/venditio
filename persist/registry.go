package persist

type Registry struct {
	tables map[string]*Table
}

func NewRegistry() *Registry {
	return &Registry{
		tables: map[string]*Table{},
	}
}

func (r *Registry) Table(name string) *Table {
	t, ok := r.tables[name]
	if !ok {
		t = NewTable(name)
		r.tables[name] = t
	}
	return t
}
