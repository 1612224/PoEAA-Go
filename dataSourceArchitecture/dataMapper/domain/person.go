package domain

import "fmt"

type Person struct {
	ID                 int
	LastName           string
	FirstName          string
	NumberOfDependents int
}

func (ps *Person) FullName() string {
	return ps.FirstName + " " + ps.LastName
}

func (ps *Person) String() string {
	return fmt.Sprintf("%d %s %d", ps.ID, ps.FullName(), ps.NumberOfDependents)
}
