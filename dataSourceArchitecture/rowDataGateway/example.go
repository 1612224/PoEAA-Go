package main

import (
	"database/sql"
	"fmt"
	"log"

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
	personOne := gateways.NewPersonGateway(-1, "Bob", "John", 1, db)
	personTwo := gateways.NewPersonGateway(-1, "Jake", "Tim", 3, db)
	personOne.Insert()
	personTwo.Insert()

	// find all
	personFinder := finders.NewPersonFinder(db)
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
	newPerson := gateways.NewPersonGateway(-1, "Huan", "Ho Minh", 0, db)
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
}
