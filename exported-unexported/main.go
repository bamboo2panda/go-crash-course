package main

import (
	"log"
	"myapp/staff"
)

var employees = []staff.Employee{
	{FirstName: "John", Lastname: "Smith", Salary: 30000, FullTime: true},
	{FirstName: "Sally", Lastname: "Jones", Salary: 50000, FullTime: true},
	{FirstName: "Mark", Lastname: "Smitters", Salary: 60000, FullTime: true},
	{FirstName: "Allan", Lastname: "Anderson", Salary: 15000, FullTime: false},
	{FirstName: "Margaret", Lastname: "Carter", Salary: 100000, FullTime: false},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}

	// log.Println(myStaff.All())
	log.Println("Overpaid staff", myStaff.Overpaid())
	log.Println("Underpaid staff", myStaff.Underpaid())
}
