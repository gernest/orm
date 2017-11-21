// Autogenerated by github.com/posener/orm
package personorm

import (
	"strings"

	"github.com/posener/orm/example"
)

const colCount = "COUNT(*)"

type PersonCount struct {
	example.Person
	Count int64
}

// String returns the SQL query string
func (s *TSelect) String() string {
	return strings.Join([]string{
		"SELECT", s.columns.String(), "FROM 'person'",
		s.where.String(),
		s.groupBy.String(),
		s.orderBy.String(),
		s.page.String(),
	}, " ")

}

// Query the database
func (s *TSelect) Query() ([]example.Person, error) {
	// create select statement
	stmt := s.String()
	args := s.Args()
	s.orm.log("Query: '%v' %v", stmt, args)
	rows, err := s.orm.db.Query(stmt, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// extract rows to structures
	var all []example.Person
	for rows.Next() {
		var item PersonCount
		if err := rows.Scan(s.scanArgs(&item)...); err != nil {
			return nil, err
		}
		all = append(all, item.Person)
	}
	return all, rows.Err()
}

// Count add a count column to the query
func (s *TSelect) Count() ([]PersonCount, error) {
	s.columns.add(colCount)
	// create select statement
	stmt := s.String()
	args := s.where.Args()
	s.orm.log("Count: '%v' %v", stmt, args)
	rows, err := s.orm.db.Query(stmt, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// extract rows to structures
	var all []PersonCount
	for rows.Next() {
		var item PersonCount
		if err := rows.Scan(s.scanArgs(&item)...); err != nil {
			return nil, err
		}
		all = append(all, item)
	}
	return all, rows.Err()
}

// SelectName Add Name to the selected column of a query
func (s *TSelect) SelectName() *TSelect {
	s.columns.add("`name`")
	return s
}

// OrderByName set order to the query results according to column name
func (s *TSelect) OrderByName(dir OrderDir) *TSelect {
	s.orderBy.add("`name`", dir)
	return s
}

// GroupByName make the query group by column name
func (s *TSelect) GroupByName() *TSelect {
	s.groupBy.add("`name`")
	return s
}

// SelectAge Add Age to the selected column of a query
func (s *TSelect) SelectAge() *TSelect {
	s.columns.add("`age`")
	return s
}

// OrderByAge set order to the query results according to column age
func (s *TSelect) OrderByAge(dir OrderDir) *TSelect {
	s.orderBy.add("`age`", dir)
	return s
}

// GroupByAge make the query group by column age
func (s *TSelect) GroupByAge() *TSelect {
	s.groupBy.add("`age`")
	return s
}

// scanArgs are list of fields to be given to the sql Scan command
func (s *TSelect) scanArgs(p *PersonCount) []interface{} {
	if len(s.columns) == 0 {
		// add to args all the fields of p
		return []interface{}{
			&p.Name,
			&p.Age,
		}
	}
	m := s.columns.indexMap()
	args := make([]interface{}, len(s.columns))
	if i := m["`name`"]; i != 0 {
		args[i-1] = &p.Name
	}
	if i := m["`age`"]; i != 0 {
		args[i-1] = &p.Age
	}
	if i := m[colCount]; i != 0 {
		args[i-1] = &p.Count
	}
	return args
}
