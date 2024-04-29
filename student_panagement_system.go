package main

import (
	"fmt"
)

var student_database = []Student{}

type Student struct {
	Id     int
	name   string
	age    int
	gender string
}

type Subject struct {
	Id    int
	title string
	mark  string
}

func menu() {
	loop := true
	var choice int
	fmt.Println("Welcome from student management system!")
	fmt.Println("=======================================")

	for loop {
		fmt.Println("What would you like to do?")
		fmt.Println("1.Add students  2.Delete students  3.Update students  4.View students")
		fmt.Scanln(&choice)
		switch choice {
		case 0:
			fmt.Println("Existing Program......")
			loop = false
		case 1:
			addStudent()
		case 4:
			viewStudent()
		default:
			/* code */
			return
		}
	}
}
func viewStudent() {
	fmt.Println(student_database)
}

func addStudent() ([]Student, error) {
	var Id int
	var name string
	var age int
	var gender string
	fmt.Println("Enter student id")
	fmt.Scanln(&Id)

	fmt.Println("Enter student name")
	fmt.Scanln(&name)

	fmt.Println("Enter student age")
	fmt.Scanln(&age)

	fmt.Println("Enter student gender")
	fmt.Scanln(&gender)

	if Id == -1 || name != "" || age != 0 || gender != "" {
		newStudent := Student{
			Id:     Id,
			name:   name,
			age:    age,
			gender: gender,
		}

		student_database = append(student_database, newStudent)
	}

	return student_database, nil

}

func main() {
	menu()
}
