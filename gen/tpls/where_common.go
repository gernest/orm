package tpls

import (
	"fmt"
	"log"
	"strings"
)

// Options are options for SQL WHERE statement
type Where struct {
	stmt []string
	args []interface{}
}

// newWhere returns a new WHERE statement
func newWhere(op Op, variable string, value interface{}) Where {
	switch op {
	case OpEq, OpNe, OpGt, OpGE, OpLt, OpLE, OpLike:
	default:
		log.Panicf("Operation %s is not defined for one value", op)
	}
	var w Where
	w.stmt = append(w.stmt, fmt.Sprintf("%s %s ?", variable, op))
	w.args = append(w.args, value)
	return w
}

// newMulWhere returns a new WHERE statement for SQL operations with more than one
// value, such as 'IN' and 'BETWEEN'.
func newMulWhere(op Op, variable string, values ...interface{}) Where {
	var w Where
	switch op {
	case OpBetween:
		if len(values) != 2 {
			log.Panicf("Operation %s accepts only 2 values, got %d values", op, len(values))
		}
		w.stmt = append(w.stmt, fmt.Sprintf("%s %s ? AND ?", variable, op))
	case OpIn:
		if len(values) > 0 {
			w.stmt = append(w.stmt, fmt.Sprintf("%s %s (%s)", variable, op, qMarks(len(values))))
		}
	default:
		log.Panicf("Operation %s does not support multiple values", op)
	}
	w.args = append(w.args, values...)
	return w
}

func (w *Where) String() string {
	if w == nil || len(w.stmt) == 0 {
		return ""
	}
	return "WHERE " + strings.Join(w.stmt, " ")
}

func (w *Where) Args() []interface{} {
	if w == nil {
		return nil
	}
	return w.args
}

func (w Where) Or(right Where) Where {
	return binary(w, right, "OR")
}

func (w Where) And(right Where) Where {
	return binary(w, right, "AND")
}

func binary(l Where, r Where, op string) Where {
	l.stmt = append([]string{"("}, l.stmt...)
	l.stmt = append(l.stmt, ")", op, "(")
	l.stmt = append(l.stmt, r.stmt...)
	l.stmt = append(l.stmt, ")")
	l.args = append(l.args, r.args...)
	return l
}