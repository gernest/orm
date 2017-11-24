// Autogenerated by github.com/posener/orm
package personorm

import (
	"database/sql"
	"database/sql/driver"
	"reflect"
	"unsafe"
)

// TSelect is the struct that holds the SELECT data
type TSelect struct {
	// Querier // TSelect is a querier, but this interface is autogenerated
	Argser
	orm *ORM
	columns
	where *Where
	groupBy
	orderBy
	page Page
}

// Rows is an alias to sql.Rows, since we need to have access to lastcols in
// scanning the results of a query, to improve the conversion performance
type Rows struct {
	*sql.Rows
}

// Values is a hack to the sql.Rows struct
// since the rows struct does not expose it's lastcols values, or a way to give
// a custom scanner to the Scan method.
// See issue https://github.com/golang/go/issues/22544
func (r *Rows) Values() []driver.Value {
	// some ugly hack to access lastcols field
	rs := reflect.ValueOf(*r)
	rs2 := reflect.New(rs.Type()).Elem()
	rs2.Set(rs)
	rf := rs2.FieldByName("lastcols")
	rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
	return rf.Interface().([]driver.Value)
}

func (s *TSelect) Args() []interface{} {
	return s.where.Args()
}

// Where applies where conditions on the query
func (s *TSelect) Where(where *Where) *TSelect {
	s.where = where
	return s
}

// Limit applies rows limit on the query response
func (s *TSelect) Limit(limit int64) *TSelect {
	s.page.limit = limit
	return s
}

// Page applies rows offset and limit on the query response
func (s *TSelect) Page(offset, limit int64) *TSelect {
	s.page.offset = offset
	s.page.limit = limit
	return s
}
