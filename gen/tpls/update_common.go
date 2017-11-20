package tpls

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

// TUpdate is a struct to hold information for an INSERT statement
type TUpdate struct {
	TInsert
	where Where
}

func (u *TUpdate) Where(where *Where) *TUpdate {
	u.where = *where
	return u
}

// Exec inserts the data to the given database
func (u *TUpdate) Exec() (sql.Result, error) {
	if len(u.cols) == 0 || len(u.values) == 0 {
		return nil, fmt.Errorf("nothing to update")
	}

	stmt := u.String()
	log.Printf("Update: '%v' (%v)", stmt, u.values)
	return u.db.Exec(stmt, append(u.values, u.where.args...)...)
}

// add adds a column and value to the UPDATE statement
func (u *TUpdate) add(name string, value interface{}) *TUpdate {
	u.TInsert.add(name, value)
	return u
}

func (u *TUpdate) assignmentList() string {
	assignments := make([]string, 0, len(u.cols))
	for _, col := range u.cols {
		assignments = append(assignments, fmt.Sprintf("%s = ?", col))
	}
	return strings.Join(assignments, ", ")
}
