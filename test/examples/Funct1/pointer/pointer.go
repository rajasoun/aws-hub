package main

import "fmt"

type Employee struct {
	Name string
	ID   string
}

func UpdateEmployeeWithPointer(Emp *Employee) {
	Emp.Name = "Rajesh"
	Emp.ID = "EMP002"
}

func UpdateEmployeeWithoutPointer(Emp Employee) {
	Emp.Name = "Rajesh"
	Emp.ID = "EMP002"
}

func main() {
	emp := Employee{Name: "Anji", ID: "EMP001"}
	var empPointer *Employee = &emp

	// call function UpdateEmployeeWithoutPointer
	UpdateEmployeeWithoutPointer(emp)
	fmt.Println(emp.Name, emp.ID)
	// output:  Anji EMP001

	// call function UpdateEmployeeWithPointer
	UpdateEmployeeWithPointer(empPointer)
	fmt.Println(emp.Name, emp.ID)
	// output:  John EMP002
}
