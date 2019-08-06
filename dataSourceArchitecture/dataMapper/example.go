package main

import (
	"database/sql"
	"fmt"
	"log"

	keygenerators "github.com/1612224/PoEAA-Go/objectRelationalStructure/identityField/keyGenerators"

	"github.com/1612224/PoEAA-Go/dataSourceArchitecture/dataMapper/mappers"
	_ "github.com/lib/pq"
)

func printAllPeople(personMapper *mappers.PersonMapper) {
	people, err := personMapper.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, ps := range people {
		fmt.Println(ps)
	}
}

func main() {
	// open db
	connStr := "user=postgres dbname=poeaa sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Open DB Successfully")

	// init db
	keyGenerator := keygenerators.NewKeyGenerator(db, "people", 1)
	personMapper := mappers.NewPersonMapper(db, keyGenerator)
	personOne := personMapper.NewPerson("Bob", "John", 1)
	personTwo := personMapper.NewPerson("Jake", "Tim", 3)
	personMapper.Insert(personOne)
	personMapper.Insert(personTwo)

	// find all
	fmt.Println("--------------Find All------------------")
	printAllPeople(personMapper)

	// find one
	fmt.Println("--------------Find One------------------")
	person, err := personMapper.FindOne(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(person)

	// insert
	fmt.Println("--------------Insert------------------")
	newPerson := personMapper.NewPerson("Huan", "Ho Minh", 12)
	if err := personMapper.Insert(newPerson); err != nil {
		log.Fatal(err)
	}
	printAllPeople(personMapper)

	// update
	fmt.Println("--------------Update------------------")
	personTwo.LastName = "Rowler"
	personTwo.NumberOfDependents = 9
	if err := personMapper.Update(personTwo); err != nil {
		log.Fatal(err)
	}
	printAllPeople(personMapper)

	// delete
	fmt.Println("--------------Delete------------------")
	if err := personMapper.Delete(newPerson); err != nil {
		log.Fatal(err)
	}
	printAllPeople(personMapper)

	// cleanup db
	personMapper.DeleteAll()
}
