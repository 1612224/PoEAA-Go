package finders

import (
	"database/sql"

	"github.com/1612224/PoEAA-Go/dataSourceArchitecture/rowDataGateway/gateways"
)

type idGenerator interface {
	NextID() (int, error)
	Reset() error
}

// PersonFinder finds person in People table
type PersonFinder struct {
	db        *sql.DB
	generator idGenerator
}

// NewPersonFinder creates new person finder
func NewPersonFinder(db *sql.DB, generator idGenerator) *PersonFinder {
	return &PersonFinder{db, generator}
}

type scanner interface {
	Scan(dest ...interface{}) error
}

func rowToPersonGateway(row scanner, db *sql.DB, generator idGenerator) (*gateways.PersonGateway, error) {
	var id int
	var lastName string
	var firstName string
	var numberOfDependents int
	err := row.Scan(&id, &lastName, &firstName, &numberOfDependents)
	if err != nil {
		return nil, err
	}
	return gateways.NewPersonGatewayWithID(id, lastName, firstName, numberOfDependents, db, generator), nil
}

// FindOne finds one person using id
func (finder *PersonFinder) FindOne(id int) (*gateways.PersonGateway, error) {
	db := finder.db
	row := db.QueryRow(`
		select id, lastname, firstname, numberofdependents
		from people
		where id = $1
	`, id)
	return rowToPersonGateway(row, db, finder.generator)
}

// FindAll finds all people
func (finder *PersonFinder) FindAll() ([]*gateways.PersonGateway, error) {
	db := finder.db
	rows, err := db.Query(`
		select id, lastname, firstname, numberofdependents
		from people
	`)

	if err != nil {
		return nil, err
	}

	result := []*gateways.PersonGateway{}
	for rows.Next() {
		gateway, err := rowToPersonGateway(rows, db, finder.generator)
		if err != nil {
			return nil, err
		}
		result = append(result, gateway)
	}
	return result, nil
}
