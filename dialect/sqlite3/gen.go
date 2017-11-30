package sqlite3

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/labstack/gommon/log"
	"github.com/posener/orm/dialect/sqltypes"
	"github.com/posener/orm/load"
)

// Gen is code generator for sqlite3 dialect
type Gen struct {
	Tp *load.Type
}

// Name returns the dialect name
func (g *Gen) Name() string {
	return "sqlite3"
}

// ConvertValueCode returns go code for converting value returned from the
// database to the given field.
func (g *Gen) ConvertValueCode(field *load.Field) string {
	s := tmpltType{
		Field:       field,
		ConvertType: g.convertType(field),
	}
	b := bytes.NewBuffer(nil)
	err := tmplt.Execute(b, s)
	if err != nil {
		log.Fatalf("executing sqlite convert value template: %s", err)
	}
	return strings.Trim(b.String(), "\n")
}

type tmpltType struct {
	ConvertType string
	Field       *load.Field
}

var tmplt = template.Must(template.New("sqlite3").Parse(`
				val, ok := vals[i].({{.ConvertType}})
				if !ok {
					return nil, fmt.Errorf(errMsg, "{{.Field.Name}}", i, vals[i], vals[i], "{{.Field.ExtTypeName}}")
				}
				tmp := {{.Field.NonPointer}}(val)
				row.{{.Field.VarName}} = {{if .Field.IsPointer -}}&{{end}}tmp
`))

// ConvertType is the type of the field when returned by sql/driver from database
func (g *Gen) convertType(f *load.Field) string {
	switch g.sqlType(f).Family() {
	case sqltypes.Integer:
		return "int64"
	case sqltypes.Float:
		return "float64"
	case sqltypes.Text, sqltypes.Blob, sqltypes.VarChar:
		return "[]byte"
	case sqltypes.Boolean:
		return "bool"
	default:
		return f.NonPointer()
	}
}

func (Gen) sqlType(f *load.Field) sqltypes.Type {
	if f.SQL.CustomType != "" {
		return f.SQL.CustomType
	}
	switch f.NonPointer() {
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return sqltypes.Integer
	case "float", "float8", "float16", "float32", "float64":
		return sqltypes.Float
	case "bool":
		return sqltypes.Boolean
	case "string":
		return sqltypes.Text
	case "[]byte":
		return sqltypes.Blob
	case "time.Time":
		return sqltypes.TimeStamp
	default:
		log.Fatalf("Unknown column type for %s", f.NonPointer())
		return sqltypes.NA
	}
}
