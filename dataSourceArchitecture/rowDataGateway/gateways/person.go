package gateways

import (
	"database/sql"
	"fmt"

	keygenerators "github.com/1612224/PoEAA-Go/objectRelationalStructure/identityField/keyGenerators"
)

type idGenerator interface {
	NextID() (int, error)
}

// PersonGateway acts as 1 row in People table
type PersonGateway struct {
	db        *sql.DB
	generator idGenerator

	ID                 int
	LastName           string
	FirstName          string
	NumberOfDependents int
}

const updateStatement = `
	update people
	set lastname = $1, firstname = $2, numberofdependents = $3
	where id = $4
`
const insertStatement = `
	insert into people (id, lastname, firstname, numberofdependents) values
	($1, $2, $3, $4)
`
const deleteStatement = `
	delete from people
	where id = $1
`

// NewPersonGatewayWithID creates new gateway with specific id
func NewPersonGatewayWithID(id int, lastName, firstName string,
	numberOfDependents int,
	db *sql.DB,
	generator idGenerator) *PersonGateway {
	return &PersonGateway{
		ID:                 id,
		FirstName:          firstName,
		LastName:           lastName,
		NumberOfDependents: numberOfDependents,
		db:                 db,
		generator:          generator,
	}
}

// NewPersonGateway creates new gateway
func NewPersonGateway(lastName, firstName string,
	numberOfDependents int,
	db *sql.DB,
	generator idGenerator) *PersonGateway {
	return NewPersonGatewayWithID(
		keygenerators.KeyPlaceholder,
		lastName, firstName,
		numberOfDependents,
		db, generator,
	)
}

// Insert inserts new row into table
func (gateway *PersonGateway) Insert() error {
	db := gateway.db

	// generate new id for new gateway if id is placeholder
	if gateway.ID == keygenerators.KeyPlaceholder {
		id, err := gateway.generator.NextID()
		if err != nil {
			return err
		}
		gateway.ID = id
	}
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
