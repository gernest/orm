package tpls

import (
	"context"
	"database/sql"

	"github.com/posener/orm/common"
	"github.com/posener/orm/dialect"
)

// DB is an interface of functions of sql.DB which are used by orm struct.
type DB interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	Close() error
}

// orm represents an orm of a given struct.
// All functions available to interact with an SQL table that is related
// to this struct, are done by an instance of this object.
// To get an instance of orm use Open or New functions.
type orm struct {
	dialect dialect.Dialect
	db      DB
	logger  Logger
}

func (o *orm) Close() error {
	return o.db.Close()
}

// Logger sets a logger to the orm package
func (o *orm) Logger(logger Logger) {
	o.logger = logger
}

// CreateBuilder builds an SQL CREATE statement parameters
type CreateBuilder struct {
	params common.CreateParams
	orm    *orm
}

// IfNotExists sets IF NOT EXISTS for the CREATE SQL statement
func (c *CreateBuilder) IfNotExists() *CreateBuilder {
	c.params.IfNotExists = true
	return c
}

// InsertBuilder builds an INSERT statement parameters
type InsertBuilder struct {
	params common.InsertParams
	orm    *orm
}

// UpdateBuilder builds SQL INSERT statement parameters
type UpdateBuilder struct {
	params common.UpdateParams
	orm    *orm
}

func (u *UpdateBuilder) Where(where common.Where) *UpdateBuilder {
	u.params.Where = where
	return u
}

// DeleteBuilder builds SQL DELETE statement parameters
type DeleteBuilder struct {
	params common.DeleteParams
	orm    *orm
}

// Where applies where conditions on the query
func (d *DeleteBuilder) Where(w common.Where) *DeleteBuilder {
	d.params.Where = w
	return d
}
