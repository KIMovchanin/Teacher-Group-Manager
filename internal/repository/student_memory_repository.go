package repository

import (
	"time"

	"github.com/KIMovchanin/Teacher-Group-Manager/internal/domain"
)

type StudentMemoryRepository struct {
	students []domain.Student
	nextID   int64
}

func NewStudentMemoryRepository() *StudentMemoryRepository {
	return &StudentMemoryRepository{
		students: []domain.Student{
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
		},
		nextID: 3,
	}
}

func (r *StudentMemoryRepository) ListStudents() []domain.Student {
	return r.students
}

func (r *StudentMemoryRepository) CreateStudent(firstName, lastName string) domain.Student {
	student := domain.Student{
		ID:        r.nextID,
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: time.Now(),
	}

	r.students = append(r.students, student)
	r.nextID++

	return student
}
