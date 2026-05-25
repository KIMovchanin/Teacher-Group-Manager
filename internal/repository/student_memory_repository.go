package repository

import (
	"sync"
	"time"

	"github.com/KIMovchanin/Teacher-Group-Manager/internal/domain"
)

type StudentMemoryRepository struct {
	mu       sync.Mutex
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
	r.mu.Lock()
	defer r.mu.Unlock()

	students := make([]domain.Student, len(r.students))
	copy(students, r.students)

	return students
}

func (r *StudentMemoryRepository) CreateStudent(firstName, lastName string) domain.Student {
	r.mu.Lock()
	defer r.mu.Unlock()

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

func (r *StudentMemoryRepository) GetStudentByID(id int64) (domain.Student, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, student := range r.students {
		if student.ID == id {
			return student, true
		}
	}

	return domain.Student{}, false
}

func (r *StudentMemoryRepository) DeleteStudent(id int64) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	for index, student := range r.students {
		if id == student.ID {
			// в кусок от students до найденного индекса добавляю всё после него
			r.students = append(r.students[:index], r.students[index+1:]...)
			return true
		}
	}
	return false
}
