package gateways

import (
	"database/sql"
)

type idGenerator interface {
	NextID() (int, error)
	Reset() error
}

// PeopleGateway acts as Gateway for People table
type PeopleGateway struct {
	db        *sql.DB
	generator idGenerator
}

// NewPeopleGateway creates new PersonGateway
func NewPeopleGateway(db *sql.DB, generator idGenerator) *PeopleGateway {
	ps := &PeopleGateway{db, generator}
	return ps
}

// FindAll returns all people
func (gateway *PeopleGateway) FindAll() (*sql.Rows, error) {
	db := gateway.db
	rows, err := db.Query("select * from people")
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// FindWithLastName return all people with specific lastname
func (gateway *PeopleGateway) FindWithLastName(lastName string) (*sql.Rows, error) {
	db := gateway.db
	rows, err := db.Query("select * from people where lastname=$1", lastName)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// FindWhere return all people with specific where clause
func (gateway *PeopleGateway) FindWhere(whereClause string) (*sql.Rows, error) {
	db := gateway.db
	rows, err := db.Query("select * from people where " + whereClause)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// FindRow return one person using id
func (gateway *PeopleGateway) FindRow(id int) *sql.Row {
	db := gateway.db
	row := db.QueryRow("select * from people where id=$1", id)
	return row
}

// Update updates
func (gateway *PeopleGateway) Update(id int, lastName, firstName string, numberOfDependents int) error {
	db := gateway.db
	_, err := db.Exec(`
		update people 
		set lastname = $1, firstname = $2, numberofdependents = $3 
		where id = $4
	`, lastName, firstName, numberOfDependents, id)
	return err
}

// Insert inserts
func (gateway *PeopleGateway) Insert(lastName, firstName string, numberOfDependents int) error {
	db := gateway.db

	id, err := gateway.generator.NextID()
	if err != nil {
		return err
	}
	_, err = db.Exec(`
		insert into people(id, lastname, firstname, numberofdependents) 
		values
			($1,$2,$3,$4)
	`, id, lastName, firstName, numberOfDependents)
	return err
}

// Delete deletes
func (gateway *PeopleGateway) Delete(id int) error {
	db := gateway.db
	_, err := db.Exec(`
		delete from people
		where id = $1
	`, id)
	return err
}

// DeleteAll deletes all
func (gateway *PeopleGateway) DeleteAll() error {
	db := gateway.db

	_, err := db.Exec(`
		delete from people
	`)
	err = gateway.generator.Reset()
	return err
}
