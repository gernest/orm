package tpls

import (
	"fmt"
	"strings"
)

// Where are options for SQL WHERE statement
type Where struct {
	stmt []string
	args []interface{}
}

// newWhere returns a new WHERE statement
func newWhere(op Op, variable string, value interface{}) *Where {
	var w Where
	w.stmt = append(w.stmt, fmt.Sprintf("%s %s ?", variable, op))
	w.args = append(w.args, value)
	return &w
}

// newWhereIn returns a new 'WHERE variable IN (...)' statement
func newWhereIn(variable string, values ...interface{}) *Where {
	var w Where
	w.stmt = append(w.stmt, fmt.Sprintf("%s IN (%s)", variable, qMarks(len(values))))
	w.args = append(w.args, values...)
	return &w
}

// newWhereBetween returns a new 'WHERE variable BETWEEN low AND high' statement
func newWhereBetween(variable string, low, high interface{}) *Where {
	var w Where
	w.stmt = append(w.stmt, fmt.Sprintf("%s BETWEEN ? AND ?", variable))
	w.args = append(w.args, low, high)
	return &w
}

// String returns the WHERE SQL statement
func (w *Where) String() string {
	if w == nil || len(w.stmt) == 0 {
		return ""
	}
	return "WHERE " + strings.Join(w.stmt, " ")
}

// Args are the arguments for executing a SELECT query with a WHERE condition
func (w *Where) Args() []interface{} {
	if w == nil {
		return nil
	}
	return w.args
}

// Or applies an or condition between two where conditions
func (w *Where) Or(right *Where) *Where {
	return binary(w, right, "OR")
}

// And applies an and condition between two where conditions
func (w *Where) And(right *Where) *Where {
	return binary(w, right, "AND")
}

func binary(l *Where, r *Where, op string) *Where {
	l.stmt = append([]string{"("}, l.stmt...)
	l.stmt = append(l.stmt, ")", op, "(")
	l.stmt = append(l.stmt, r.stmt...)
	l.stmt = append(l.stmt, ")")
	l.args = append(l.args, r.args...)
	return l
}
