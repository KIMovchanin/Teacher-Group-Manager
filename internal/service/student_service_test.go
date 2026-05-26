package service

import (
	"testing"
	"time"

	"github.com/KIMovchanin/Teacher-Group-Manager/internal/domain"
)

// unit test
// Test with fake dependency - лучше, когда есть зависимотсти у тестируемого объекта

type fakeStudentRepository struct {
	createStudentCalled bool
	createdStudent      domain.Student
}

func (r *fakeStudentRepository) ListStudents() []domain.Student {
	return nil
}

func (r *fakeStudentRepository) CreateStudent(firstName, lastName string) domain.Student {
	r.createStudentCalled = true

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

func TestStudentService_CreateStudent(t *testing.T) {
	tests := []struct {
		name                 string
		firstName            string
		lastName             string
		wantErr              bool
		wantRepositoryCalled bool
	}{
		{
			name:                 "valid student",
			firstName:            "Ivan",
			lastName:             "Petrov",
			wantErr:              false,
			wantRepositoryCalled: true,
		},
		{
			name:                 "empty first name",
			firstName:            "",
			lastName:             "Petrov",
			wantErr:              true,
			wantRepositoryCalled: false,
		},
		{
			name:                 "empty last name",
			firstName:            "Ivan",
			lastName:             "",
			wantErr:              true,
			wantRepositoryCalled: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &fakeStudentRepository{}
			studentService := NewStudentService(repo)

			student, err := studentService.CreateStudent(tt.firstName, tt.lastName)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}

				if repo.createStudentCalled {
					t.Fatal("expected repository not to be called")
				}

				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if !repo.createStudentCalled {
				t.Fatal("expected repository to be called")
			}

			if student.FirstName != tt.firstName {
				t.Fatalf("expected first name %s, got %s", tt.firstName, student.FirstName)
			}

			if student.LastName != tt.lastName {
				t.Fatalf("expected last name %s, got %s", tt.lastName, student.LastName)
			}
		})
	}
}
