package mappers

import (
	"database/sql"

	keygenerators "github.com/1612224/PoEAA-Go/objectRelationalStructure/identityField/keyGenerators"

	"github.com/1612224/PoEAA-Go/dataSourceArchitecture/dataMapper/domain"
)

type idGenerator interface {
	NextID() (int, error)
	Reset() error
}

type PersonMapper struct {
	db        *sql.DB
	generator idGenerator
}

func NewPersonMapper(db *sql.DB, generator idGenerator) *PersonMapper {
	return &PersonMapper{db, generator}
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
		select id, lastname, firstname, numberofdependents
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
		select id, lastname, firstname, numberofdependents
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
	if person.ID == keygenerators.KeyPlaceholder {
		id, err := pm.generator.NextID()
		if err != nil {
			return err
		}
		person.ID = id
	}
	_, err := db.Exec(`
		insert into people (id, lastname, firstname, numberofdependents)
		values
			($1, $2, $3, $4)
	`, person.ID, person.LastName, person.FirstName, person.NumberOfDependents)
	return err
}

func (pm *PersonMapper) Update(person *domain.Person) error {
	db := pm.db
	_, err := db.Exec(`
		update people 
		set lastname = $1, firstname = $2, numberofdependents = $3
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
	if err != nil {
		return err
	}
	err = pm.generator.Reset()
	return err
}

func (pm *PersonMapper) NewPerson(lastname, firstname string, numberOfDependents int) *domain.Person {
	person := &domain.Person{
		LastName:           lastname,
		FirstName:          firstname,
		NumberOfDependents: numberOfDependents,
	}
	person.ID = keygenerators.KeyPlaceholder
	return person
}
