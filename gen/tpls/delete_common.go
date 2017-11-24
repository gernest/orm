package tpls

import (
	"github.com/posener/orm"
)

// Delete is the struct that holds the SELECT data
type Delete struct {
	orm.Delete
	orm *ORM
}

// Where applies where conditions on the query
func (d *Delete) Where(w orm.Where) *Delete {
	d.Delete.Where = w
	return d
}
