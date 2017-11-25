// Autogenerated by github.com/posener/orm; DO NOT EDIT
package allsqlite3

import (
	"database/sql"

	"github.com/posener/orm/common"
)

const table = "all"

// Open opens database connection
func Open(dataSourceName string) (*ORM, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &ORM{db: db}, nil
}

// Select returns an object to create a SELECT statement
func (o *ORM) Select() *Select {
	s := &Select{
		internal: common.Select{
			Table: table,
		},
		orm: o,
	}
	s.internal.Columns = &s.columns
	return s
}

// Insert returns a new INSERT statement
func (o *ORM) Insert() *Insert {
	return &Insert{
		internal: common.Insert{Table: table},
		orm:      o,
	}
}

// Update returns a new UPDATE statement
func (o *ORM) Update() *Update {
	return &Update{
		internal: common.Update{Table: table},
		orm:      o,
	}
}

// Delete returns an object for a DELETE statement
func (o *ORM) Delete() *Delete {
	return &Delete{
		internal: common.Delete{Table: table},
		orm:      o,
	}
}