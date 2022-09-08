package school

import (
	"sort"
)

// type School
// func New() *School
// func (s *School) Add(string, int)
// func (s *School) Grade(int) []string
// func (s *School) Enrollment() []Grade

type Grade struct {
	grade int
	names []string
}

type School struct {
	grades []Grade
}

func New() *School {
	school := School{}
	return &school
}

func (s *School) Add(name string, grade int) {
	success := false
	for index, value := range s.grades {
		if value.grade == grade {
			s.grades[index].names = append(s.grades[index].names, name)
			success = true
		}
	}
	if success == false {
		names := []string{name}
		temp := Grade{grade, names}
		s.grades = append(s.grades, temp)
	}
}

func (s *School) Grade(grade int) []string {
	for _, value := range s.grades {
		if value.grade == grade {
			return value.names
		}
	}
	return []string{}
}

func (s *School) Enrollment() []Grade {
	allData := s.grades
	sort.Slice(allData, func(i, j int) bool {
		return (allData[i].grade) < (allData[j].grade)
	})

	for _, v := range allData {
		sort.Strings(v.names)
	}
	return allData

}
