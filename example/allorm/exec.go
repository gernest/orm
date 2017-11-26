// Autogenerated by github.com/posener/orm; DO NOT EDIT
package allorm

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/posener/orm/example"
)

// Exec creates a table for the given struct
func (b *CreateBuilder) Exec() (sql.Result, error) {
	stmt, args := b.orm.dialect.Create(&b.params)
	b.orm.log("Create: '%v' %v", stmt, args)
	return b.orm.db.ExecContext(contextOrBackground(b.params.Ctx), stmt, args...)
}

// query is used by the Select.Query and Select.Limit functions
func (b *SelectBuilder) query(ctx context.Context) (*sql.Rows, error) {
	stmt, args := b.orm.dialect.Select(&b.params)
	b.orm.log("Query: '%v' %v", stmt, args)
	return b.orm.db.QueryContext(ctx, stmt, args...)
}

// Exec inserts the data to the given database
func (b *InsertBuilder) Exec() (sql.Result, error) {
	if len(b.params.Assignments) == 0 {
		return nil, fmt.Errorf("nothing to insert")
	}
	stmt, args := b.orm.dialect.Insert(&b.params)
	b.orm.log("Insert: '%v' %v", stmt, args)
	return b.orm.db.ExecContext(contextOrBackground(b.params.Ctx), stmt, args...)
}

// Exec inserts the data to the given database
func (b *UpdateBuilder) Exec() (sql.Result, error) {
	if len(b.params.Assignments) == 0 {
		return nil, fmt.Errorf("nothing to update")
	}
	stmt, args := b.orm.dialect.Update(&b.params)
	b.orm.log("Update: '%v' %v", stmt, args)
	return b.orm.db.ExecContext(contextOrBackground(b.params.Ctx), stmt, args...)
}

// Exec runs the delete statement on a given database.
func (b *DeleteBuilder) Exec() (sql.Result, error) {
	stmt, args := b.orm.dialect.Delete(&b.params)
	b.orm.log("Delete: '%v' %v", stmt, args)
	return b.orm.db.ExecContext(contextOrBackground(b.params.Ctx), stmt, args...)
}

// Query the database
func (b *SelectBuilder) Query() ([]example.All, error) {
	ctx := contextOrBackground(b.params.Ctx)
	rows, err := b.query(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// extract rows to structures
	var all []example.All
	for rows.Next() {
		// check context cancellation
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		item, err := scan(b.orm.dialect.Name(), b.columns, rows)
		if err != nil {
			return nil, err
		}
		all = append(all, item.All)
	}
	return all, rows.Err()
}

// Count add a count column to the query
func (b *SelectBuilder) Count() ([]AllCount, error) {
	ctx := contextOrBackground(b.params.Ctx)
	b.columns.count = true
	rows, err := b.query(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// extract rows to structures
	var all []AllCount
	for rows.Next() {
		// check context cancellation
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		item, err := scan(b.orm.dialect.Name(), b.columns, rows)
		if err != nil {
			return nil, err
		}
		all = append(all, *item)
	}
	return all, rows.Err()
}

func contextOrBackground(ctx context.Context) context.Context {
	if ctx == nil {
		return context.Background()
	}
	return ctx
}
