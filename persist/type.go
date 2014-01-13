package persist

import (
	"database/sql"
)

const (
	DRIVER_SQLITE   = "sqlite3"
	DRIVER_MYSQL    = "mysql"
	DRIVER_POSTGRES = "postgres"
)

type Typeable interface {
	RawType() interface{}
}

type Column struct {
	Unique        bool
	PrimaryKey    bool
	AutoIncrement bool
	Unsigned      bool
	Type          Typeable
}

type Bool struct {
	Typeable
}

func (t *Bool) RawType() interface{} {
	return &sql.NullBool{}
}

type SmallInt struct {
	Typeable
}

func (t *SmallInt) RawType() interface{} {
	return &sql.NullInt64{}
}

type Integer struct {
	Typeable
}

func (t *Integer) RawType() interface{} {
	return &sql.NullInt64{}
}

type BigInt struct {
	Typeable
}

func (t *BigInt) RawType() interface{} {
	return &sql.NullInt64{}
}

type Decimal struct {
	Typeable
}

func (t *Decimal) RawType() interface{} {
	return &sql.NullFloat64{}
}

type Float struct {
	Typeable
}

func (t *Float) RawType() interface{} {
	return &sql.NullFloat64{}
}

type Double struct {
	Typeable
}

func (t *Double) RawType() interface{} {
	return &sql.NullFloat64{}
}

type String struct {
	Typeable
}

func (t *String) RawType() interface{} {
	return &sql.NullString{}
}

type Text struct {
	Typeable
}

func (t *Text) RawType() interface{} {
	return &sql.NullString{}
}
