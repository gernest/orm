// Autogenerated by github.com/posener/orm
package allorm

import "database/sql"

type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

func New(db DB) *ORM {
	return &ORM{db: db}
}

type ORM struct {
	db DB
}

// Create returns a struct for a CREATE statement
func (o *ORM) Create() *TCreate {
	return &TCreate{db: o.db}
}

// Select returns an object to create a SELECT statement
func (o *ORM) Select() *TSelect {
	return &TSelect{db: o.db}
}

// Insert returns a new INSERT statement
func (o *ORM) Insert() *TInsert {
	return &TInsert{db: o.db}
}

// Insert returns a new INSERT statement
func (o *ORM) Update() *TUpdate {
	return &TUpdate{TInsert: TInsert{db: o.db}}
}

// Delete returns an object for a DELETE statement
func (o *ORM) Delete() *TDelete {
	return &TDelete{db: o.db}
}
