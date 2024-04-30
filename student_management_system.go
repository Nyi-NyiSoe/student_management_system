package main

import (
	"errors"
	"fmt"
	"strconv"
)

var student_database = []Student{}

type Student struct {
	Id       string
	name     string
	age      string
	gender   string
	Subjects []Subject
}

type Subject struct {
	title string
	mark  string
}

func menu() {
	var choice int

	fmt.Println("Welcome from student management system!")
	fmt.Println("=======================================")

	for {
		fmt.Println("What would you like to do?")
		fmt.Println("1.Add students  2.Delete students  3.Update students  4.View students  5.Assign Subject")
		fmt.Scanln(&choice)
		switch choice {
		case 0:
			fmt.Println("Exiting program")
			return
		case 1:
			addStudent()
		case 3:
			updateStudent()
		case 4:
			viewStudent()
		case 5:
			assignSubject()
		default:
			fmt.Println("Exiting program")
			return

		}

	}
}

func updateStudent() {
	var Id int
	var newEditValue string
	for {
		fmt.Println("Enter the ID of student")
		_, err := fmt.Scanln(&Id)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		break
	}
	for index, value := range student_database {
		oldId, err := strconv.Atoi(value.Id)
		if err != nil {

			return
		}
		if oldId == Id {
			fmt.Println("Which field do you want do edit?")
			var choice string
			fmt.Scanln(&choice)
			switch choice {
			case "id":
				fmt.Println("Enter new Id")
				fmt.Scanln(&newEditValue)
				student_database[index].Id = newEditValue
			case "name":
				fmt.Println("Enter new name")
				fmt.Scanln(&newEditValue)
				student_database[index].name = newEditValue
			case "age":
				fmt.Println("Enter new age")
				fmt.Scanln(&newEditValue)
				student_database[index].age = newEditValue
			case "gender":
				fmt.Println("Enter new gender")
				fmt.Scanln(&newEditValue)
				student_database[index].gender = newEditValue
			case "Subjects":
				assignSubject()
			}
		} else {
			fmt.Println("Student doesnt exist")
			continue
		}
	}

}

func assignSubject() {
	var Id int
	var title string
	var mark string
	for {
		fmt.Println("Enter the ID of student")
		_, err := fmt.Scanln(&Id)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		break
	}

	for index, value := range student_database {
		oldValue, err := strconv.Atoi(value.Id)
		if err != nil {
			fmt.Println("Student doesnt exist")
			return
		}
		if Id == oldValue {
			fmt.Println("Enter subject title")
			fmt.Scanln(&title)
			fmt.Println("Enter marks ")
			fmt.Scanln(&mark)

			newSubjects := Subject{
				title: title,
				mark:  mark,
			}
			student_database[index].Subjects = append(student_database[index].Subjects, newSubjects)
			break
		}

	}

}
func viewStudent() {
	fmt.Println("||\tID\t||\tName\t||\tAge\t||\tGender\t||\tSubjects\t")
	fmt.Println("=================================================================================")
	fmt.Println("||\t\t||\t\t||\t\t||\t\t||")
	for i := 0; i < len(student_database); i++ {
		var subjects []string
		for j := 0; j < len(student_database[i].Subjects); j++ {
			subjects = append(subjects, student_database[i].Subjects[j].title)
		}
		fmt.Printf("||\t%s\t||\t%s\t||\t%s\t||\t%s\t||\t%s\t\n", student_database[i].Id, student_database[i].name, student_database[i].age, student_database[i].gender, subjects)
		fmt.Println("||\t\t||\t\t||\t\t||\t\t||")
	}

}
func addStudent() ([]Student, error) {
	var Id string
	var name string
	var age string
	var gender string

	for {

		fmt.Println("Enter student's id")
		_, err := fmt.Scanln(&Id)
		if err != nil {
			fmt.Println("Invalid input.")
			continue
		}

		var idExist bool = false
		for _, student := range student_database {
			if student.Id == Id {
				fmt.Println("Student ID already exists")
				idExist = true
				break
			}

		}
		if idExist {
			continue
		}

		break

	}
	for {

		fmt.Println("Enter student's name")
		_, err := fmt.Scanln(&name)
		if err != nil {
			fmt.Println("Invalid input.")
			continue
		}
		break

	}
	for {

		fmt.Println("Enter student's age")
		_, err := fmt.Scanln(&age)
		if err != nil {
			fmt.Println("Invalid input.")
			continue
		}
		break

	}
	for {

		fmt.Println("Enter student's gender")
		_, err := fmt.Scanln(&gender)
		if err != nil {
			fmt.Println("Invalid input.")
			continue
		}
		break

	}

	if Id != "" || name != "" || age != "" || gender != "" {
		newStudent := Student{
			Id:       Id,
			name:     name,
			age:      age,
			gender:   gender,
			Subjects: []Subject{},
		}
		student_database = append(student_database, newStudent)
	} else {
		return nil, errors.New("invalid input")
	}

	return student_database, nil

}

func main() {
	menu()

}
