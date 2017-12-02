// Package personorm was auto-generated by github.com/posener/orm; DO NOT EDIT
package personorm

import (
	"context"
	"database/sql/driver"

	"github.com/posener/orm/common"
	"github.com/posener/orm/example"
)

type Scanner interface {
	Columns() []string
	First(dialect string, values []driver.Value) (*example.Person, error)
}

// PersonCount is a struct for counting rows of type Person
type PersonCount struct {
	example.Person
	Count int64
}

// SelectBuilder builds an SQL SELECT statement parameters
type SelectBuilder struct {
	params   common.SelectParams
	conn     *conn
	selector selector
}

func (b *SelectBuilder) Scanner() Scanner {
	return b.params.Columns.(*selector)
}

// Where applies where conditions on the query
func (b *SelectBuilder) Where(where common.Where) *SelectBuilder {
	b.params.Where = where
	return b
}

// Limit applies rows limit on the query response
func (b *SelectBuilder) Limit(limit int64) *SelectBuilder {
	b.params.Page.Limit = limit
	return b
}

// Page applies rows offset and limit on the query response
func (b *SelectBuilder) Page(offset, limit int64) *SelectBuilder {
	b.params.Page.Offset = offset
	b.params.Page.Limit = limit
	return b
}

// SelectName adds Name to the selected column of a query
func (b *SelectBuilder) SelectName() *SelectBuilder {
	b.selector.SelectName = true
	return b
}

// OrderByName set order to the query results according to column name
func (b *SelectBuilder) OrderByName(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("name", dir)
	return b
}

// GroupByName make the query group by column name
func (b *SelectBuilder) GroupByName() *SelectBuilder {
	b.params.Groups.Add("name")
	return b
}

// SelectAge adds Age to the selected column of a query
func (b *SelectBuilder) SelectAge() *SelectBuilder {
	b.selector.SelectAge = true
	return b
}

// OrderByAge set order to the query results according to column age
func (b *SelectBuilder) OrderByAge(dir common.OrderDir) *SelectBuilder {
	b.params.Orders.Add("age", dir)
	return b
}

// GroupByAge make the query group by column age
func (b *SelectBuilder) GroupByAge() *SelectBuilder {
	b.params.Groups.Add("age")
	return b
}

// Context sets the context for the SQL query
func (b *SelectBuilder) Context(ctx context.Context) *SelectBuilder {
	b.params.Ctx = ctx
	return b
}
