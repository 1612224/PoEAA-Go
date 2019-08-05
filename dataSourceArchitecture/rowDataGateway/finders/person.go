package finders

import (
	"database/sql"

	"github.com/1612224/PoEAA-Go/dataSourceArchitecture/rowDataGateway/gateways"
)

// PersonFinder finds person in People table
type PersonFinder struct {
	db *sql.DB
}

// NewPersonFinder creates new person finder
func NewPersonFinder(db *sql.DB) *PersonFinder {
	return &PersonFinder{db}
}

type scanner interface {
	Scan(dest ...interface{}) error
}

func rowToPersonGateway(row scanner, db *sql.DB) (*gateways.PersonGateway, error) {
	var id int
	var lastName string
	var firstName string
	var numberOfDependents int
	err := row.Scan(&id, &lastName, &firstName, &numberOfDependents)
	if err != nil {
		return nil, err
	}
	return gateways.NewPersonGateway(id, lastName, firstName, numberOfDependents, db), nil
}

// FindOne finds one person using id
func (finder *PersonFinder) FindOne(id int) (*gateways.PersonGateway, error) {
	db := finder.db
	row := db.QueryRow(`
		select id, lastname, firstname, number_of_dependents
		from people
		where id = $1
	`, id)
	return rowToPersonGateway(row, db)
}

// FindAll finds all people
func (finder *PersonFinder) FindAll() ([]*gateways.PersonGateway, error) {
	db := finder.db
	rows, err := db.Query(`
		select id, lastname, firstname, number_of_dependents
		from people
	`)

	if err != nil {
		return nil, err
	}

	result := []*gateways.PersonGateway{}
	for rows.Next() {
		gateway, err := rowToPersonGateway(rows, db)
		if err != nil {
			return nil, err
		}
		result = append(result, gateway)
	}
	return result, nil
}
