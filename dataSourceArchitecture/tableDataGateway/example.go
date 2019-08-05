package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/1612224/PoEAA-Go/dataSourceArchitecture/tableDataGateway/gateways"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres dbname=poeaa sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Open DB Successfully")
	personGateway := gateways.NewPersonGateway(db)

	rows, err := personGateway.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var firstName string
		var lastName string
		var numberOfDependents int
		rows.Scan(&id, &lastName, &firstName, &numberOfDependents)
		fmt.Println(id, lastName, firstName, numberOfDependents)
	}
}
