package service

import (
	"errors"

	"github.com/KIMovchanin/Teacher-Group-Manager/internal/domain"
	"github.com/KIMovchanin/Teacher-Group-Manager/internal/repository"
)

type StudentService struct {
	studentRepository *repository.StudentMemoryRepository
}

func NewStudentService(studentRepository *repository.StudentMemoryRepository) *StudentService {
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
