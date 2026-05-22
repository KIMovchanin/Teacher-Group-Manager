package service

import (
	"errors"

	"github.com/KIMovchanin/Teacher-Group-Manager/internal/domain"
)

type StudentRepository interface {
	ListStudents() []domain.Student
	CreateStudent(firstName, lastName string) domain.Student
	GetStudentByID(id int64) (domain.Student, bool)
	DeleteStudent(id int64) bool
}

type StudentService struct {
	studentRepository StudentRepository
}

func NewStudentService(studentRepository StudentRepository) *StudentService {
	return &StudentService{
		studentRepository: studentRepository,
	}
}

func (s *StudentService) ListStudents() []domain.Student {
	return s.studentRepository.ListStudents()
}

func (s *StudentService) CreateStudent(firstName, lastName string) (domain.Student, error) {
	if firstName == "" || lastName == "" {
		return domain.Student{}, errors.New("first name and last name are required")
	}

	return s.studentRepository.CreateStudent(firstName, lastName), nil
}

func (s *StudentService) GetStudentByID(id int64) (domain.Student, error) {
	student, found := s.studentRepository.GetStudentByID(id)
	if !found {
		return domain.Student{}, errors.New("student not found")
	}
	return student, nil
}

func (s *StudentService) DeleteStudent(id int64) error {
	deleted := s.studentRepository.DeleteStudent(id)
	if !deleted {
		return errors.New("student not found")
	}
	return nil
}
