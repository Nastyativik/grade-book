package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	students := []*Student{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Журнал")
		fmt.Println("1. Добавить студента")
		fmt.Println("2. Добавить оценку")
		fmt.Println("3. Показать всех студентов")
		fmt.Println("4. Удалить студента")
		fmt.Println("5. Показать студентов со средним баллом < 4.0")
		fmt.Println("6. Выйти")
		fmt.Print("Ваш выбор: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			students = addStudent(scanner, students)
		case "2":
			students = addGrade(scanner, students)
		case "3":
			showAllStudents(students)
		case "4":
			students = deleteStudent(scanner, students)
		case "5":
			filterByLowAverage(students)
		case "6":
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Пожалуйста, введите номер от 1 до 6.")
		}
	}
}
