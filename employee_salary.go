package main

import (
	"fmt"
	"reflect"
)

type salary struct {
	Salary float32 `json:"salary"`
}

func NewSalary() *salary {
	s := new(salary)
	s.Salary = 1000.0
	return s
}

func (s *salary) getRaise(percent float32) {
	s.Salary = s.Salary * (1. + percent/100)
}

func main() {
	s := NewSalary()
	fmt.Printf("salary is %f$\n", s.Salary)
	fmt.Println("I'm get raised")
	s.getRaise(20)
	fmt.Printf("salary is %f$\n", s.Salary)
	fmt.Println(reflect.TypeOf(s))
}
