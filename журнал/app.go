package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func addStudent(scanner *bufio.Scanner, students []*Student) []*Student {
	fmt.Print("Фамилия: ")
	scanner.Scan()
	lastName := strings.TrimSpace(scanner.Text())

	fmt.Print("Имя: ")
	scanner.Scan()
	firstName := strings.TrimSpace(scanner.Text())

	fmt.Print("Отчество: ")
	scanner.Scan()
	middleName := strings.TrimSpace(scanner.Text())

	student := &Student{
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
		Grades:     []int{},
	}

	students = append(students, student)
	fmt.Println("Студент", lastName, firstName, middleName, "добавлен.")
	return students
}

func addGrade(scanner *bufio.Scanner, students []*Student) []*Student {
	if len(students) == 0 {
		fmt.Println("Нет студентов в журнале.")
		return students
	}

	fmt.Println("Выберите студента:")
	for i, s := range students {
		fmt.Println(i+1, ".", s.LastName, s.FirstName, s.MiddleName)
	}

	fmt.Print("Номер студента: ")
	scanner.Scan()
	num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil || num < 1 || num > len(students) {
		fmt.Println("Неверный номер.")
		return students
	}

	student := students[num-1]

	fmt.Print("Оценка (2–5): ")
	scanner.Scan()
	gradeStr := strings.TrimSpace(scanner.Text())
	grade, err := strconv.Atoi(gradeStr)
	if err != nil || grade < 2 || grade > 5 {
		fmt.Println("Оценка должна быть от 2 до 5.")
		return students
	}

	student.AddGrade(grade)
	fmt.Println("Оценка", grade, "добавлена студенту", student.LastName, student.FirstName, student.MiddleName)
	return students
}

func showAllStudents(students []*Student) {
	if len(students) == 0 {
		fmt.Println("Журнал пуст.")
		return
	}
	fmt.Println("Журнал студентов")
	for _, s := range students {
		s.DisplayInfo()
	}
}

func filterByLowAverage(students []*Student) {
	found := false
	fmt.Println("Студенты со средним баллом ниже 4.0:")
	for _, s := range students {
		avg := s.CalculateAverage()
		if avg < 4.0 {
			avgStr := strconv.FormatFloat(avg, 'f', 2, 64)
			fmt.Println("- ", s.LastName, s.FirstName, s.MiddleName, ":", avgStr)
			found = true
		}
	}
	if !found {
		fmt.Println("Нет таких студентов.")
	}
}

func deleteStudent(scanner *bufio.Scanner, students []*Student) []*Student {
	if len(students) == 0 {
		fmt.Println("Журнал пуст.")
		return students
	}

	fmt.Println("Выберите студента для удаления:")
	for i, s := range students {
		fmt.Println(i+1, ".", s.LastName, s.FirstName, s.MiddleName)
	}

	fmt.Print("Номер студента: ")
	scanner.Scan()
	num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil || num < 1 || num > len(students) {
		fmt.Println("Неверный номер.")
		return students
	}

	removed := students[num-1]
	fmt.Println("Студент", removed.LastName, removed.FirstName, removed.MiddleName, "удалён.")

	return append(students[:num-1], students[num:]...)
}
