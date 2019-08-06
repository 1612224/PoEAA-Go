package main

import (
	"database/sql"
	"fmt"
	"log"

	keygenerators "github.com/1612224/PoEAA-Go/objectRelationalStructure/identityField/keyGenerators"

	"github.com/1612224/PoEAA-Go/dataSourceArchitecture/rowDataGateway/finders"
	"github.com/1612224/PoEAA-Go/dataSourceArchitecture/rowDataGateway/gateways"
	_ "github.com/lib/pq"
)

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
	personOne := gateways.NewPersonGateway("Bob", "John", 1, db, keyGenerator)
	personTwo := gateways.NewPersonGateway("Jake", "Tim", 3, db, keyGenerator)
	personOne.Insert()
	personTwo.Insert()

	// find all
	fmt.Println("----------------Find all---------------------")
	personFinder := finders.NewPersonFinder(db, keyGenerator)
	allPeople, _ := personFinder.FindAll()
	for _, gw := range allPeople {
		fmt.Println(gw)
	}

	// find one
	fmt.Println("----------------Find one---------------------")
	person, _ := personFinder.FindOne(1)
	fmt.Println(person)

	// insert and print
	fmt.Println("-----------------Insert----------------------")
	newPerson := gateways.NewPersonGateway("Huan", "Ho Minh", 0, db, keyGenerator)
	err = newPerson.Insert()
	if err != nil {
		log.Fatal(err)
	}
	allPeople, _ = personFinder.FindAll()
	for _, gw := range allPeople {
		fmt.Println(gw)
	}

	// update and print
	fmt.Println("-----------------Update----------------------")
	personTwo.LastName = "Rick"
	personTwo.NumberOfDependents = 5
	err = personTwo.Update()
	if err != nil {
		log.Fatal(err)
	}
	allPeople, _ = personFinder.FindAll()
	for _, gw := range allPeople {
		fmt.Println(gw)
	}

	// delete and print
	fmt.Println("---------------------Delete------------------")
	err = newPerson.Delete()
	if err != nil {
		log.Fatal(err)
	}
	allPeople, _ = personFinder.FindAll()
	for _, gw := range allPeople {
		fmt.Println(gw)
	}

	// cleanup db
	for _, gw := range allPeople {
		gw.Delete()
	}
	if err := keyGenerator.Reset(); err != nil {
		log.Fatal(err)
	}
}
