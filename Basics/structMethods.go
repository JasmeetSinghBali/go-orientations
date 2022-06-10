package main

import (
	"fmt"
)

type Student struct {
	name  string
	marks []int
	age   int
}

// getAge () method that act on student object s of type Student struct & returns int
func (s Student) getAge() int {
	return s.age
}

// setAge () method that act on student object s of type *Student pointer struct & returns int
func (s *Student) setAge(age int) {
	s.age = age
}

// getavgGrades ()
func (s *Student) getavgMarks() float32 {
	sum := 0
	for _, m := range s.marks {
		sum += m
	}
	return float32(sum / len(s.marks))
}

//getMaxMark()
func (s *Student) getMaxMark() int {
	currMax := 0
	for _, m := range s.marks {
		if m > currMax {
			currMax = m
		}
	}
	return currMax
}

func main() {
	s1 := &Student{"Tim", []int{20, 30, 40, 50}, 14}
	fmt.Println("this is an object of type struct Student:/n", *s1)
	ageGrabbed := s1.getAge()
	fmt.Println("grab age of the Student object age by getAge() method:\n", ageGrabbed)
	s1.setAge(24)
	fmt.Println("age updated via setAge() method:\n", s1.age)
	avg := s1.getavgMarks()
	fmt.Println("avg marks getavgMarks() method:\n", avg)
	s2 := &Student{"Jas", []int{100, 0, 0, 50}, 5}
	avgs2 := s2.getavgMarks()
	fmt.Println("This is new student: \n", s2)
	fmt.Println("avg marks for Jas: via getavgMarks() \n", avgs2)
	maxMark := s2.getMaxMark()
	fmt.Println("max mark scored by Jas: via getMaxMark() \n", maxMark)
}
