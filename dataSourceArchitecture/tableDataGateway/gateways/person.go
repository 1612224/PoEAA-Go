package gateways

import "database/sql"

// PersonGateway acts as Gateway for People table
type PersonGateway struct {
	db *sql.DB
}

// NewPersonGateway creates new PersonGateway
func NewPersonGateway(db *sql.DB) *PersonGateway {
	return &PersonGateway{db}
}

// FindAll returns all people
func (personGateway *PersonGateway) FindAll() (*sql.Rows, error) {
	db := personGateway.db
	rows, err := db.Query("select * from people")
	if err != nil {
		return nil, err
	}
	return rows, nil
}
