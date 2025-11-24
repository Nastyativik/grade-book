package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	FirstName  string
	MiddleName string
	LastName   string
	Grades     []int
}

func (s *Student) CalculateAverage() float64 {
	if len(s.Grades) == 0 {
		return 0.0
	}
	sum := 0
	for _, grade := range s.Grades {
		sum += grade
	}
	return float64(sum) / float64(len(s.Grades))
}

func (s *Student) AddGrade(grade int) {
	s.Grades = append(s.Grades, grade)
}

func (s *Student) DisplayInfo() {
	avg := s.CalculateAverage()
	avgStr := strconv.FormatFloat(avg, 'f', 2, 64)
	fmt.Println("ФИО:", s.LastName, s.FirstName, s.MiddleName)
	fmt.Println("Оценки:", s.Grades)
	fmt.Println("Средний балл:", avgStr)
	fmt.Println()
}
