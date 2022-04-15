package main

import "fmt"

type address struct {
	id    string
	line1 string
	line2 string
}

// custom data type similar like class. but no concept of class
type employee struct {
	id                string
	name              string
	department        string
	isManager         bool
	skills            []string
	permanentAddress  address
	reseidenceAddress address
}

/*
	Method. A method is a function with receiver argument.
	With the help of receiver argument we can define method on any type.
	In below example : updateDepartment() method is defined on employee type

	similarity :: this is similar concept like we define method in class

	Java :
	class Employee {
		Employee(){...}

		public void updateDepartment(String newDepartment) {...}
	}


	Python :
	class Employee:
		def __init__(self):
			....

		def updateDepartment(self, newDepartment):
			....
*/
func (e *employee) updateDepartment(newDepartment string) {
	e.department = newDepartment
}

func (e *employee) addSkill(newSkill string) {
	e.skills = append(e.skills, newSkill)
}

func main() {
	e1 := employee{}
	e1.updateDepartment("android")
	e1.addSkill("golang")
	fmt.Printf("%+v", e1)
}

/**
	go to next example 09-value-and-pointer-receiver
**/
