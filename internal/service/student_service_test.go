package service

import (
	"testing"
	"time"

	"github.com/KIMovchanin/Teacher-Group-Manager/internal/domain"
)

// unit test
// Test with fake dependency - лучше, когда есть зависимотсти у тестируемого объекта

type fakeStudentRepository struct {
	createdStudent domain.Student
}

func (r *fakeStudentRepository) ListStudents() []domain.Student {
	return nil
}

func (r *fakeStudentRepository) CreateStudent(firstName, lastName string) domain.Student {
	r.createdStudent = domain.Student{
		ID:        1,
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: time.Now(),
	}

	return r.createdStudent
}

func (r *fakeStudentRepository) GetStudentByID(id int64) (domain.Student, bool) {
	return domain.Student{}, false
}

func (r *fakeStudentRepository) DeleteStudent(id int64) bool {
	return false
}

func TestStudentService_CreateStudent_ReturnsErrorWhenFirstNameIsEmpty(t *testing.T) {
	repo := &fakeStudentRepository{}
	service := NewStudentService(repo)

	_, err := service.CreateStudent("", "Petrov")

	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestStudentService_CreatedStudent_CreatesStudent(t *testing.T) {
	repo := &fakeStudentRepository{}
	service := NewStudentService(repo)

	student, err := service.CreateStudent("Ivan", "Petrov")
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if student.FirstName != "Ivan" {
		t.Fatalf("expected first name Ivan, got %s", student.FirstName)
	}

	if student.LastName != "Petrov" {
		t.Fatalf("expected last name Petrov, got %s", student.LastName)
	}

	if repo.createdStudent.ID != student.ID {
		t.Fatalf("expected repository to create student with id %d, got %d", student.ID, repo.createdStudent.ID)
	}
}
