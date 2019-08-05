package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/1612224/PoEAA-Go/dataSourceArchitecture/tableDataGateway/gateways"
	_ "github.com/lib/pq"
)

// Scanner scans
type Scanner interface {
	Scan(dest ...interface{}) error
}

func printPerson(scanner Scanner) {
	var id int
	var firstName string
	var lastName string
	var numberOfDependents int
	if err := scanner.Scan(&id, &lastName, &firstName, &numberOfDependents); err != nil {
		log.Print(err)
	} else {
		fmt.Println(id, lastName, firstName, numberOfDependents)
	}
}

func main() {
	connStr := "user=postgres dbname=poeaa sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Open DB Successfully")
	peopleGateway := gateways.NewPeopleGateway(db)

	// init db
	peopleGateway.Insert("Bob", "John", 1)
	peopleGateway.Insert("Jake", "Tim", 3)

	// print all
	fmt.Println("---------------Find all---------------------")
	rows, err := peopleGateway.FindAll()
	// rows, err := peopleGateway.FindWithLastName("Jack")
	// rows, err := peopleGateway.FindWhere("number_of_dependents <= 2")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		printPerson(rows)
	}

	// print one
	fmt.Println("----------------Find one---------------------")
	row := peopleGateway.FindRow(2)
	printPerson(row)

	// insert and print
	fmt.Println("--------------Insert------------------------")
	err = peopleGateway.Insert("Huan", "Ho Minh", 0)
	rows, err = peopleGateway.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		printPerson(rows)
	}

	// update and print
	fmt.Println("--------------Update-------------------------")
	err = peopleGateway.Update(1, "Tim", "Jake", 4)
	if err != nil {
		log.Fatal(err)
	}
	rows, err = peopleGateway.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		printPerson(rows)
	}

	// delete and print
	fmt.Println("--------------Delete---------------------------")
	err = peopleGateway.Delete(3)
	if err != nil {
		log.Fatal(err)
	}
	rows, err = peopleGateway.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		printPerson(rows)
	}

	// clean up db
	peopleGateway.DeleteAll()
}
