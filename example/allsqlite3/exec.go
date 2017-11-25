// Autogenerated by github.com/posener/orm; DO NOT EDIT
package allsqlite3

import (
	"database/sql"
	"fmt"

	"github.com/posener/orm/dialect/sqlite3"
)

const createString = `CREATE TABLE 'all' ( 'auto' INTEGER PRIMARY KEY AUTOINCREMENT, 'notnil' TEXT NOT NULL, 'int' INTEGER, 'int8' INTEGER, 'int16' INTEGER, 'int32' INTEGER, 'int64' INTEGER, 'uint' INTEGER, 'uint8' INTEGER, 'uint16' INTEGER, 'uint32' INTEGER, 'uint64' INTEGER, 'time' TIMESTAMP, 'varcharstring' VARCHAR(100), 'varcharbyte' VARCHAR(100), 'string' TEXT, 'bytes' BLOB, 'bool' BOOLEAN, 'pint' INTEGER, 'pint8' INTEGER, 'pint16' INTEGER, 'pint32' INTEGER, 'pint64' INTEGER, 'puint' INTEGER, 'puint8' INTEGER, 'puint16' INTEGER, 'puint32' INTEGER, 'puint64' INTEGER, 'ptime' TIMESTAMP, 'pvarcharstring' VARCHAR(100), 'pvarcharbyte' VARCHAR(100), 'pstring' TEXT, 'pbytes' BLOB, 'pbool' BOOLEAN, 'select' INTEGER )`

// Exec creates a table for the given struct
func (c *Create) Exec() (sql.Result, error) {
	c.orm.log("Create: '%v'", createString)
	return c.orm.db.Exec(createString)
}

func (s *Select) query() (*sql.Rows, error) {
	stmt, args := sqlite3.Select(&s.internal)
	s.orm.log("Query: '%v' %v", stmt, args)
	return s.orm.db.Query(stmt, args...)
}

// Exec inserts the data to the given database
func (i *Insert) Exec() (sql.Result, error) {
	if len(i.internal.Assignments) == 0 {
		return nil, fmt.Errorf("nothing to insert")
	}
	stmt, args := sqlite3.Insert(&i.internal)
	i.orm.log("Insert: '%v' %v", stmt, args)
	return i.orm.db.Exec(stmt, args...)
}

// Exec inserts the data to the given database
func (u *Update) Exec() (sql.Result, error) {
	if len(u.internal.Assignments) == 0 {
		return nil, fmt.Errorf("nothing to update")
	}
	stmt, args := sqlite3.Update(&u.internal)
	u.orm.log("Update: '%v' %v", stmt, args)
	return u.orm.db.Exec(stmt, args...)
}

// Exec runs the delete statement on a given database.
func (d *Delete) Exec() (sql.Result, error) {
	stmt, args := sqlite3.Delete(&d.internal)
	d.orm.log("Delete: '%v' %v", stmt, args)
	return d.orm.db.Exec(stmt, args...)
}
