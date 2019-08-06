package mappers

import (
	"database/sql"

	"github.com/1612224/PoEAA-Go/dataSourceArchitecture/dataMapper/domain"
)

var currentID = 0

func getNextID() int {
	currentID++
	return currentID
}

type PersonMapper struct {
	db *sql.DB
}

func NewPersonMapper(db *sql.DB) *PersonMapper {
	return &PersonMapper{db}
}

type scanner interface {
	Scan(dest ...interface{}) error
}

func rowToPerson(row scanner) (*domain.Person, error) {
	var id int
	var lastName, firstName string
	var numberOfDependents int
	err := row.Scan(&id, &lastName, &firstName, &numberOfDependents)
	if err != nil {
		return nil, err
	}
	return &domain.Person{
		ID:                 id,
		LastName:           lastName,
		FirstName:          firstName,
		NumberOfDependents: numberOfDependents,
	}, nil
}

func (pm *PersonMapper) FindAll() ([]*domain.Person, error) {
	db := pm.db
	rows, err := db.Query(`
		select id, lastname, firstname, number_of_dependents
		from people
	`)
	if err != nil {
		return nil, err
	}

	result := []*domain.Person{}
	for rows.Next() {
		person, err := rowToPerson(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, person)
	}
	return result, nil
}

func (pm *PersonMapper) FindOne(id int) (*domain.Person, error) {
	db := pm.db
	row := db.QueryRow(`
		select id, lastname, firstname, number_of_dependents
		from people
		where id = $1
	`, id)

	person, err := rowToPerson(row)
	if err != nil {
		return nil, err
	}
	return person, nil
}

func (pm *PersonMapper) Insert(person *domain.Person) error {
	db := pm.db
	_, err := db.Exec(`
		insert into people (id, lastname, firstname, number_of_dependents)
		values
			($1, $2, $3, $4)
	`, person.ID, person.LastName, person.FirstName, person.NumberOfDependents)
	return err
}

func (pm *PersonMapper) Update(person *domain.Person) error {
	db := pm.db
	_, err := db.Exec(`
		update people 
		set lastname = $1, firstname = $2, number_of_dependents = $3
		where id = $4
	`, person.LastName, person.FirstName, person.NumberOfDependents, person.ID)
	return err
}

func (pm *PersonMapper) Delete(person *domain.Person) error {
	db := pm.db
	_, err := db.Exec(`
		delete from people
		where id = $1
	`, person.ID)
	return err
}

func (pm *PersonMapper) DeleteAll() error {
	db := pm.db
	_, err := db.Exec(`delete from people`)
	return err
}

func (pm *PersonMapper) NewPerson(lastname, firstname string, numberOfDependents int) *domain.Person {
	person := &domain.Person{
		LastName:           lastname,
		FirstName:          firstname,
		NumberOfDependents: numberOfDependents,
	}
	person.ID = getNextID()
	return person
}
