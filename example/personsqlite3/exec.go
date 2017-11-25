// Autogenerated by github.com/posener/orm; DO NOT EDIT
package personsqlite3

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/posener/orm/dialect/sqlite3"
)

const createString = `CREATE TABLE 'person' ( 'name' TEXT, 'age' INTEGER )`

// Exec creates a table for the given struct
func (c *Create) Exec(ctx context.Context) (sql.Result, error) {
	c.orm.log("Create: '%v'", createString)
	return c.orm.db.ExecContext(ctx, createString)
}

func (s *Select) query(ctx context.Context) (*sql.Rows, error) {
	stmt, args := sqlite3.Select(&s.internal)
	s.orm.log("Query: '%v' %v", stmt, args)
	return s.orm.db.QueryContext(ctx, stmt, args...)
}

// Exec inserts the data to the given database
func (i *Insert) Exec(ctx context.Context) (sql.Result, error) {
	if len(i.internal.Assignments) == 0 {
		return nil, fmt.Errorf("nothing to insert")
	}
	stmt, args := sqlite3.Insert(&i.internal)
	i.orm.log("Insert: '%v' %v", stmt, args)
	return i.orm.db.ExecContext(ctx, stmt, args...)
}

// Exec inserts the data to the given database
func (u *Update) Exec(ctx context.Context) (sql.Result, error) {
	if len(u.internal.Assignments) == 0 {
		return nil, fmt.Errorf("nothing to update")
	}
	stmt, args := sqlite3.Update(&u.internal)
	u.orm.log("Update: '%v' %v", stmt, args)
	return u.orm.db.ExecContext(ctx, stmt, args...)
}

// Exec runs the delete statement on a given database.
func (d *Delete) Exec(ctx context.Context) (sql.Result, error) {
	stmt, args := sqlite3.Delete(&d.internal)
	d.orm.log("Delete: '%v' %v", stmt, args)
	return d.orm.db.ExecContext(ctx, stmt, args...)
}
