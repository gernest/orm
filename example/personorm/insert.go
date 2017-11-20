// Autogenerated by github.com/posener/orm
package personorm

import (
	"fmt"
	"strings"

	"github.com/posener/orm/example"
)

func (i *TInsert) String() string {
	return fmt.Sprintf(`INSERT INTO person (%s) VALUES (%s)`,
		strings.Join(i.cols, ", "),
		qMarks(len(i.values)),
	)
}

// InsertPerson creates an INSERT statement according to the given object
func (o *ORM) InsertPerson(p *example.Person) *TInsert {
	i := o.Insert()
	i.add("name", p.Name)
	i.add("age", p.Age)
	return i
}

// SetName sets value for column name in the INSERT statement
func (i *TInsert) SetName(value string) *TInsert {
	return i.add("name", value)
}

// SetAge sets value for column age in the INSERT statement
func (i *TInsert) SetAge(value int) *TInsert {
	return i.add("age", value)
}
