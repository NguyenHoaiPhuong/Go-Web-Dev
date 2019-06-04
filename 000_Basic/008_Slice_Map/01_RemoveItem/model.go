package main

// StudentID : globalstudent ID
var StudentID int

// Student : struct
type Student struct {
	id      int
	spammer []int
}

// University : struct
type University struct {
	students   []*Student
	studentmap map[int]*Student
}

// NewStudent : Create new student
func NewStudent(uni *University) *Student {
	StudentID++
	s := new(Student)
	s.id = StudentID
	s.spammer = make([]int, 131072) // 1 MiB
	s.Register(uni)

	return s
}

// Init : initialize the university
func (uni *University) Init() {
	uni.students = make([]*Student, 0)
	uni.studentmap = make(map[int]*Student)
}

// Register : register student to the University
func (s *Student) Register(uni *University) {
	uni.students = append(uni.students, s)
	uni.studentmap[s.id] = s
}

// RemoveStudentFromSlice : Remove Student From Slice
func (uni *University) RemoveStudentFromSlice(studentID int) {
	idx := -1
	for i, s := range uni.students {
		if s.id == studentID {
			idx = i
			break
		}
	}
	if idx == -1 {
		return
	}

	var newStudents []*Student
	if numOfStudents := len(uni.students); numOfStudents > 0 {
		// Special handling of last element
		if idx == numOfStudents-1 {
			newStudents = uni.students[:numOfStudents-1]
		} else {
			newStudents = append(uni.students[:idx], uni.students[idx+1:]...)
		}
	}
	uni.students = newStudents
}

// RemoveStudentFromMap : Remove Student From Map
func (uni *University) RemoveStudentFromMap(studentID int) {
	delete(uni.studentmap, studentID)
}
