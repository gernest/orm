package tpls

type Query struct {
	sel   Select
	where Where
}

func NewQuery() Query {
	return Query{}
}

func (q Query) Select(s Select) Query {
	q.sel = s
	return q
}

func (q Query) Where(w Where) Query {
	q.where = w
	return q
}