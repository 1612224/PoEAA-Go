package gateways

import (
	"database/sql"
	"fmt"
)

var currentID = 0

func getNextID() int {
	currentID++
	return currentID
}

// PersonGateway acts as 1 row in People table
type PersonGateway struct {
	db *sql.DB

	ID                 int
	LastName           string
	FirstName          string
	NumberOfDependents int
}

const updateStatement = `
	update people
	set lastname = $1, firstname = $2, number_of_dependents = $3
	where id = $4
`
const insertStatement = `
	insert into people (id, lastname, firstname, number_of_dependents) values
	($1, $2, $3, $4)
`
const deleteStatement = `
	delete from people
	where id = $1
`

// NewPersonGateway creates new gateway
func NewPersonGateway(id int, lastName, firstName string, numberOfDependents int, db *sql.DB) *PersonGateway {
	finalID := id
	if id == -1 {
		finalID = getNextID()
	}

	return &PersonGateway{
		ID:                 finalID,
		FirstName:          firstName,
		LastName:           lastName,
		NumberOfDependents: numberOfDependents,
		db:                 db,
	}
}

// Insert inserts new row into table
func (gateway *PersonGateway) Insert() error {
	db := gateway.db
	_, err := db.Exec(
		insertStatement,
		gateway.ID,
		gateway.LastName,
		gateway.FirstName,
		gateway.NumberOfDependents,
	)
	return err
}

// Update updates existing row in table
func (gateway *PersonGateway) Update() error {
	db := gateway.db
	_, err := db.Exec(
		updateStatement,
		gateway.LastName,
		gateway.FirstName,
		gateway.NumberOfDependents,
		gateway.ID,
	)
	return err
}

// Delete deletes existing row in table
func (gateway *PersonGateway) Delete() error {
	db := gateway.db
	_, err := db.Exec(
		deleteStatement,
		gateway.ID,
	)
	return err
}

func (gateway *PersonGateway) String() string {
	return fmt.Sprintf("%d, %s, %s, %d",
		gateway.ID,
		gateway.LastName,
		gateway.FirstName,
		gateway.NumberOfDependents,
	)
}
