package persist

type Registry struct {
	Tables map[string]*Table
}

func NewRegistry() *Registry {
	return &Registry{
		Tables: map[string]*Table{},
	}
}

func (r *Registry) Table(name string) *Table {
	t, ok := r.Tables[name]
	if !ok {
		t = NewTable(name)
		r.Tables[name] = t
	}
	return t
}
