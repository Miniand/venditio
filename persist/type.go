package persist

import (
	"strings"
)

const (
	DRIVER_SQLITE   = "sqlite3"
	DRIVER_MYSQL    = "mysql"
	DRIVER_POSTGRES = "postgres"
)

type Typeable interface {
	String(driver string) string
}

type BaseType struct {
	Index         bool
	Unique        bool
	PrimaryKey    bool
	AutoIncrement bool
	Unsigned      bool
}

func (t *BaseType) Suffix(driver string) string {
	parts := []string{}
	if t.Unsigned {

	}
	return strings.Join(parts, " ")
}

type SmallInt struct {
	BaseType
	Typeable
}

func (t *SmallInt) String(driver string) string {
	var col string
	switch driver {
	case DRIVER_SQLITE:
		col = "INTEGER(2)"
	case DRIVER_MYSQL:
		col = "SMALLINT"
	case DRIVER_POSTGRES:
		col = "smallint"
	default:
		panic("Unknown driver")
	}
	return col + t.Suffix(driver)
}

type Integer struct {
	BaseType
	Typeable
}

func (t *Integer) String(driver string) string {
	var col string
	switch driver {
	case DRIVER_SQLITE:
		col = "INTEGER(4)"
	case DRIVER_MYSQL:
		col = "INT"
	case DRIVER_POSTGRES:
		col = "integer"
	default:
		panic("Unknown driver")
	}
	return col + t.Suffix(driver)
}

type BigInt struct {
	BaseType
	Typeable
}

func (t *BigInt) String(driver string) string {
	var col string
	switch driver {
	case DRIVER_SQLITE:
		col = "INTEGER(4)"
	case DRIVER_MYSQL:
		col = "INT"
	case DRIVER_POSTGRES:
		col = "integer"
	default:
		panic("Unknown driver")
	}
	return col + t.Suffix(driver)
}

type Decimal struct {
	BaseType
	Typeable
}

func (t *Decimal) String(driver string) string {
	var col string
	switch driver {
	case DRIVER_SQLITE:
		col = "REAL"
	case DRIVER_MYSQL:
		col = "DECIMAL"
	case DRIVER_POSTGRES:
		col = "decimal"
	default:
		panic("Unknown driver")
	}
	return col + t.Suffix(driver)
}

type Float struct {
	BaseType
	Typeable
}

func (t *Float) String(driver string) string {
	var col string
	switch driver {
	case DRIVER_SQLITE:
		col = "REAL"
	case DRIVER_MYSQL:
		col = "FLOAT"
	case DRIVER_POSTGRES:
		col = "real"
	default:
		panic("Unknown driver")
	}
	return col + t.Suffix(driver)
}

type Double struct {
	BaseType
	Typeable
}

func (t *Double) String(driver string) string {
	var col string
	switch driver {
	case DRIVER_SQLITE:
		col = "REAL"
	case DRIVER_MYSQL:
		col = "DOUBLE"
	case DRIVER_POSTGRES:
		col = "double precision"
	default:
		panic("Unknown driver")
	}
	return col + t.Suffix(driver)
}

type String struct {
	BaseType
	Typeable
}

func (t *String) String(driver string) string {
	var col string
	switch driver {
	case DRIVER_SQLITE:
		col = "TEXT"
	case DRIVER_MYSQL:
		col = "VARCHAR"
	case DRIVER_POSTGRES:
		col = "varchar"
	default:
		panic("Unknown driver")
	}
	return col + t.Suffix(driver)
}

type Text struct {
	BaseType
	Typeable
}

func (t *Text) String(driver string) string {
	var col string
	switch driver {
	case DRIVER_SQLITE:
		col = "TEXT"
	case DRIVER_MYSQL:
		col = "TEXT"
	case DRIVER_POSTGRES:
		col = "text"
	default:
		panic("Unknown driver")
	}
	return col + t.Suffix(driver)
}
