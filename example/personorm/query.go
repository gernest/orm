package personorm

import (
	"database/sql"
	"log"

	"github.com/posener/orm/where"

	"github.com/posener/orm/example"
)

type Query struct {
	sel   *Select
	where where.Options
}

func NewQuery() Query {
	return Query{}
}

func (q Query) Select(s Select) Query {
	q.sel = &s
	return q
}

func (q Query) Where(w where.Options) Query {
	q.where = w
	return q
}

func (q *Query) String() string {
	return "SELECT " + q.sel.String() + " FROM person " + q.where.String()
}

func (q *Query) Exec(db *sql.DB) ([]example.Person, error) {
	// create select statement
	stmt := q.String()
	args := q.where.Args()
	log.Printf("Query: '%v' %v", stmt, args)
	rows, err := db.Query(stmt, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// extract rows to structures
	var all []example.Person
	for rows.Next() {
		var i example.Person
		if err := rows.Scan(q.sel.scanArgs(&i)...); err != nil {
			return nil, err
		}
		all = append(all, i)
	}
	return all, rows.Err()
}
