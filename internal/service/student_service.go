package service

import (
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

func (s *StudentService) CreateStudent(firstName, lastName string) domain.Student {
	return domain.Student{
		ID:        3,
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: time.Now(),
	}
}
