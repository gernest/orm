package dialect

import (
	"fmt"
	"strings"

	"github.com/posener/orm/dialect/mysql"
	"github.com/posener/orm/dialect/sqlite3"
	"github.com/posener/orm/dialect/sqltypes"
	"github.com/posener/orm/graph"
	"github.com/posener/orm/load"
)

// Generator is API for different dialects
type Generator interface {
	// Name is the dialect name
	Name() string
	// ColumnsStatement returns the fields parts of SQL CREATE TABLE statement
	// for a specific struct and specific dialect.
	// It is used by the generation tool.
	ColumnsStatement(gr *graph.Graph) string
	// ConvertValueCode returns go code for converting value returned from the
	// database to the given field.
	ConvertValueCode(tp *load.Type, field *load.Field) string
}

// NewGen returns all known Generators
func NewGen() []Generator {
	return []Generator{
		&gen{GenImplementer: new(mysql.Gen)},
		&gen{GenImplementer: new(sqlite3.Gen)},
	}
}

type gen struct {
	GenImplementer
}

type GenImplementer interface {
	Name() string
	GoTypeToColumnType(*load.Type) sqltypes.Type
	ColumnCreateString(string, *load.Field, sqltypes.Type) string
	ConvertValueCode(*load.Type, *load.Field, sqltypes.Type) string
}

// ColumnsStatement returns the fields parts of SQL CREATE TABLE statement
func (g *gen) ColumnsStatement(gr *graph.Graph) string {
	var (
		colStmts []string
		fkStmts  []string
	)
	for _, f := range gr.Fields {
		if !f.IsReference() {
			colStmts = append(colStmts, g.ColumnCreateString(f.Columns()[0], f, g.columnType(f, 0)))
		}
	}

	// define foreign keys for the outgoing references
	for _, e := range gr.Out {
		eColStmts, eFKStmts := g.foreignKeys(e)
		colStmts = append(colStmts, eColStmts...)
		fkStmts = append(fkStmts, eFKStmts...)
	}
	stmts := append(colStmts, fkStmts...)
	return strings.Join(stmts, ", ")
}

func (g *gen) ConvertValueCode(tp *load.Type, field *load.Field) string {
	return g.GenImplementer.ConvertValueCode(tp, field, g.columnType(field, 0))
}

func (g *gen) columnType(field *load.Field, i int) sqltypes.Type {
	if custom := field.CustomType; custom != "" {
		return custom
	}
	return g.GoTypeToColumnType(field.SetTypes()[i])
}

func (g *gen) foreignKeys(outEdge graph.Edge) (colStmts []string, fkStmts []string) {
	cols := outEdge.SrcField.Columns()
	dstFields := outEdge.RelationType().PrimaryKeys
	for i := range cols {
		colStmts = append(colStmts,
			fmt.Sprintf("`%s` %s", cols[i], g.GoTypeToColumnType(&dstFields[i].Type)))
		fkStmts = append(fkStmts,
			fmt.Sprintf("FOREIGN KEY (`%s`) REFERENCES `%s`(`%s`)",
				cols[i], outEdge.RelationType().Table(), dstFields[i].Columns()[0]))
	}
	return
}
