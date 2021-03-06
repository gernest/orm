package sqltypes

import (
	"fmt"
	"regexp"
	"strconv"
)

var typeFormat = regexp.MustCompile(`([^(]+)(\((\d+)\))?`)

// Type represents an SQL column type
type Type struct {
	Name string
	Size int
}

func New(s string) (*Type, error) {
	t := new(Type)
	m := typeFormat.FindStringSubmatch(s)
	switch len(m) {
	case 0:
		return nil, fmt.Errorf("invalid SQL type: %s", s)
	case 1:
		t.Name = m[0]
	case 4:
		t.Name = m[1]
		t.Size, _ = strconv.Atoi(m[3])
	}
	return t, nil
}

func (t *Type) String() string {
	if t.Size == 0 {
		return t.Name
	}
	return fmt.Sprintf("%s(%d)", t.Name, t.Size)
}

// List of SQL types
const (
	Integer   = "INTEGER"
	Float     = "FLOAT"
	Boolean   = "BOOLEAN"
	Text      = "TEXT"
	Blob      = "BLOB"
	TimeStamp = "TIMESTAMP"
	DateTime  = "DATETIME"
	VarChar   = "VARCHAR"
)
