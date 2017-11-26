package {{.Package}}

import (
	"context"
	"database/sql"
	"fmt"

    "{{.Type.ImportPath}}"
)

// Exec creates a table for the given struct
func (c *CreateBuilder) Exec(ctx context.Context) (sql.Result, error) {
	stmt, args := c.orm.dialect.Create(&c.params)
	c.orm.log("Create: '%v' %v", stmt, args)
	return c.orm.db.ExecContext(ctx, stmt, args...)
}

// query is used by the Select.Query and Select.Limit functions
func (s *SelectBuilder) query(ctx context.Context) (*sql.Rows, error) {
	stmt, args := s.orm.dialect.Select(&s.params)
	s.orm.log("Query: '%v' %v", stmt, args)
	return s.orm.db.QueryContext(ctx, stmt, args...)
}

// Exec inserts the data to the given database
func (i *InsertBuilder) Exec(ctx context.Context) (sql.Result, error) {
	if len(i.params.Assignments) == 0 {
		return nil, fmt.Errorf("nothing to insert")
	}
	stmt, args := i.orm.dialect.Insert(&i.params)
	i.orm.log("Insert: '%v' %v", stmt, args)
	return i.orm.db.ExecContext(ctx, stmt, args...)
}

// Exec inserts the data to the given database
func (u *UpdateBuilder) Exec(ctx context.Context) (sql.Result, error) {
	if len(u.params.Assignments) == 0 {
		return nil, fmt.Errorf("nothing to update")
	}
	stmt, args := u.orm.dialect.Update(&u.params)
	u.orm.log("Update: '%v' %v", stmt, args)
	return u.orm.db.ExecContext(ctx, stmt, args...)
}

// Exec runs the delete statement on a given database.
func (d *DeleteBuilder) Exec(ctx context.Context) (sql.Result, error) {
	stmt, args := d.orm.dialect.Delete(&d.params)
	d.orm.log("Delete: '%v' %v", stmt, args)
	return d.orm.db.ExecContext(ctx, stmt, args...)
}

// Query the database
func (s *SelectBuilder) Query(ctx context.Context) ([]{{.Type.FullName}}, error) {
    rows, err := s.query(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// extract rows to structures
	var all []{{.Type.FullName}}
	for rows.Next() {
	    // check context cancellation
	    if err := ctx.Err(); err != nil  {
	        return nil, err
	    }
		item, err := scan(s.orm.dialect.Name(), s.columns, rows)
        if err != nil {
			return nil, err
		}
		all = append(all, item.{{.Type.Name}})
	}
	return all, rows.Err()
}

// Count add a count column to the query
func (s *SelectBuilder) Count(ctx context.Context) ([]{{.Type.Name}}Count, error) {
    s.columns.count = true
    rows, err := s.query(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// extract rows to structures
	var all []{{.Type.Name}}Count
	for rows.Next() {
	    // check context cancellation
	    if err := ctx.Err(); err != nil  {
	        return nil, err
	    }
		item, err := scan(s.orm.dialect.Name(), s.columns, rows)
        if err != nil {
			return nil, err
		}
		all = append(all, *item)
	}
	return all, rows.Err()
}

