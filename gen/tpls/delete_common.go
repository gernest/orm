package tpls

import (
	"database/sql"
	"fmt"
	"log"
)

// Delete returns an object for a DELETE statement
func Delete() *TDelete {
	return &TDelete{}
}

// Select is the struct that holds the SELECT data
type TDelete struct {
	fmt.Stringer
	where *Where
}

// Where applies where conditions on the query
func (d *TDelete) Where(w *Where) *TDelete {
	d.where = w
	return d
}

// Exec runs the delete statement on a given database.
func (d *TDelete) Exec(db SQLExecer) (sql.Result, error) {
	// create select statement
	stmt := d.String()
	args := d.where.Args()
	log.Printf("Delete: '%v' %v", stmt, args)
	return db.Exec(stmt, args...)
}
