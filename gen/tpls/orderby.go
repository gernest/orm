package tpls

import (
	"fmt"
	"strings"
)

type OrderDir string

const (
	Asc  OrderDir = "ASC"
	Desc OrderDir = "DESC"
)

type orderBy []string

// String is the SQL representation of GROUP BY
func (g orderBy) String() string {
	if len(g) == 0 {
		return ""
	}
	return "ORDER BY " + strings.Join(g, ", ")
}

// Add adds a column to the grouping
func (g *orderBy) add(column string, dir OrderDir) {
	*g = append(*g, fmt.Sprintf("%s %s", column, dir))
}