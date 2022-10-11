package school

import(
    "sort"
)
// Define the Grade and School types here.
type Grade struct{
    class int
    students []string
}
type School struct{
    maxGrade int 
    roster map[int]*Grade
}
func New() *School {
	return &School{ roster : make(map[int]*Grade)}
}
func (s *School) Add(student string, grade int) {
	_, ok := s.roster[grade]
	if grade > s.maxGrade {
		s.maxGrade = grade
	}
	if !ok {
		s.roster[grade] = &Grade{class: grade, students: []string{student}}
	} else {
		s.roster[grade].students = append(s.roster[grade].students, student)
	}
} 
func (s *School) Grade(level int) []string {
    _, ok := s.roster[level]
	if !ok {
		return []string{}
	}
	return s.roster[level].students
}
func (s *School) Enrollment() []Grade {
	allStudents := []Grade{}
	for i := 1; i <= s.maxGrade; i++ {
        _,ok := s.roster[i]
        if(ok){
            sort.Strings(s.roster[i].students)
			allStudents = append(allStudents, *s.roster[i])
        }
	}
	return allStudents
}
