package tpls

import "database/sql"

func Open(driverName, dataSourceName string) (*ORM, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &ORM{
		db: db,
	}, nil
}

type ORM struct {
	db     *sql.DB
	logger Logger
}

func (o *ORM) Close() error {
	return o.db.Close()
}

// Create returns a struct for a CREATE statement
func (o *ORM) Create() *TCreate {
	return &TCreate{orm: o}
}

// Insert returns a new INSERT statement
func (o *ORM) Insert() *TInsert {
	return &TInsert{orm: o}
}

// Insert returns a new INSERT statement
func (o *ORM) Update() *TUpdate {
	return &TUpdate{TInsert: TInsert{orm: o}}
}

// Delete returns an object for a DELETE statement
func (o *ORM) Delete() *TDelete {
	return &TDelete{orm: o}
}

// Logger sets a logger to the ORM package
func (o *ORM) Logger(logger Logger) {
	o.logger = logger
}
