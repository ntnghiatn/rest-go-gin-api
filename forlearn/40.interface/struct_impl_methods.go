package main

import "fmt"

type Employee interface {
	GetName() string
	GetSalary(int32) float32
}

type EmployeeImpl struct {
	Name   string
	Salary float32
}

// GetName implements Employee.
func (ec *EmployeeImpl) GetName() string {
	// panic("unimplemented")
	return ec.Name
}

// Salary implements Employee.
func (*EmployeeImpl) GetSalary(int32) float32 {
	panic("unimplemented")
}

func NewEmployee() Employee {
	return &EmployeeImpl{Name: "Kha", Salary: 10000}
}

func main() {
	e := NewEmployee()
	fmt.Println(e)
	empName := e.GetName()
	fmt.Println(empName)
}
