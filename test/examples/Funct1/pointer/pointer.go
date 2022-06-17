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

	// here calling the function UpdateEmployeeWithoutPointer
	UpdateEmployeeWithoutPointer(emp)
	fmt.Println(emp.Name, emp.ID)
	// output:  Anji EMP001

	// here calling the function UpdateEmployeeWithPointer
	UpdateEmployeeWithPointer(empPointer)
	fmt.Println(emp.Name, emp.ID)
	// output:  Rajesh EMP002
}
