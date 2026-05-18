package service

import (
	"errors"
	"time"

	"github.com/KIMovchanin/Teacher-Group-Manager/internal/domain"
)

type StudentService struct {
}

func NewStudentService() *StudentService {
	return &StudentService{}
}

func (s *StudentService) ListStudents() []domain.Student {
	return []domain.Student{
		{
			ID:        1,
			FirstName: "Ivan",
			LastName:  "Petrov",
			CreatedAt: time.Now(),
		},
		{
			ID:        2,
			FirstName: "Anna",
			LastName:  "Sidorova",
			CreatedAt: time.Now(),
		},
	}
}

func (s *StudentService) CreateStudent(firstName, lastName string) (domain.Student, error) {
	if firstName == "" || lastName == "" {
		return domain.Student{}, errors.New("first name and last name are required")
	}

	return domain.Student{
		ID:        3,
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: time.Now(),
	}, nil
}
