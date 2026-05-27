package service

import (
	"testing"
	"time"

	"github.com/KIMovchanin/Teacher-Group-Manager/internal/domain"
)

// unit test
// Test with fake dependency - лучше, когда есть зависимотсти у тестируемого объекта

type fakeStudentRepository struct {
	// CreateStudent
	createStudentCalled bool
	createdStudent      domain.Student

	// GetStudentByID
	studentToReturn domain.Student
	studentFound    bool

	// DeleteStudent
	deleteStudentCalled bool
	deleteStudentResult bool

	// ListStudent
	listStudentCalled bool
	studentsToReturn  []domain.Student
}

func (r *fakeStudentRepository) ListStudents() []domain.Student {
	r.listStudentCalled = true

	return r.studentsToReturn
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
	return r.studentToReturn, r.studentFound
}

func (r *fakeStudentRepository) DeleteStudent(id int64) bool {
	r.deleteStudentCalled = true

	return r.deleteStudentResult
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

			if repo.createStudentCalled != tt.wantRepositoryCalled {
				t.Fatalf("expected repository called %v, got %v", tt.wantRepositoryCalled, repo.createStudentCalled)
			}

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}

				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
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

func TestStudentService_GetStudentByID(t *testing.T) {
	tests := []struct {
		name        string
		id          int64
		repoStudent domain.Student
		repoFound   bool
		wantErr     bool
		wantStudent domain.Student
	}{
		{
			name: "student found",
			id:   1,
			repoStudent: domain.Student{
				ID:        1,
				FirstName: "Ivan",
				LastName:  "Petrov",
			},
			repoFound: true,
			wantErr:   false,
			wantStudent: domain.Student{
				ID:        1,
				FirstName: "Ivan",
				LastName:  "Petrov",
			},
		},
		{
			name:        "student not found",
			id:          999,
			repoStudent: domain.Student{},
			repoFound:   false,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &fakeStudentRepository{
				studentToReturn: tt.repoStudent,
				studentFound:    tt.repoFound,
			}
			studentService := NewStudentService(repo)

			student, err := studentService.GetStudentByID(tt.id)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if student.ID != tt.wantStudent.ID {
				t.Fatalf("expected student id %d, got %d", tt.wantStudent.ID, student.ID)
			}

			if student.FirstName != tt.wantStudent.FirstName {
				t.Fatalf("expected first name %s, got %s", tt.wantStudent.FirstName, student.FirstName)
			}

			if student.LastName != tt.wantStudent.LastName {
				t.Fatalf("expected last name %s, got %s", tt.wantStudent.LastName, student.LastName)
			}

		})
	}
}

func TestStudentService_DeleteStudent(t *testing.T) {
	tests := []struct {
		name                 string
		id                   int64
		repoDeleted          bool
		wantErr              bool
		wantRepositoryCalled bool
	}{
		{
			name:                 "student deleted",
			id:                   1,
			repoDeleted:          true,
			wantErr:              false,
			wantRepositoryCalled: true,
		},
		{
			name:                 "student not found",
			id:                   999,
			repoDeleted:          false,
			wantErr:              true,
			wantRepositoryCalled: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &fakeStudentRepository{
				deleteStudentResult: tt.repoDeleted,
			}
			studentService := NewStudentService(repo)

			err := studentService.DeleteStudent(tt.id)

			if repo.deleteStudentCalled != tt.wantRepositoryCalled {
				t.Fatalf("expected repository called %v, got %v", tt.wantRepositoryCalled, repo.deleteStudentCalled)
			}

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
		})
	}

}

func TestStudentService_ListStudents(t *testing.T) {
	repo := &fakeStudentRepository{
		studentsToReturn: []domain.Student{
			{
				ID:        1,
				FirstName: "Ivan",
				LastName:  "Petrov",
			},
			{
				ID:        2,
				FirstName: "Anna",
				LastName:  "Sidorova",
			},
		},
	}

	studentService := NewStudentService(repo)

	students := studentService.ListStudents()

	if !repo.listStudentCalled {
		t.Fatal("expected repository ListStudents to be called")
	}

	if len(students) != 2 {
		t.Fatalf("expected 2 students, got %d", len(students))
	}

	if students[0].FirstName != "Ivan" {
		t.Fatalf("expected first student Ivan, got %s", students[0].FirstName)
	}

	if students[1].FirstName != "Anna" {
		t.Fatalf("expected second student Anna, got %s", students[1].FirstName)
	}
}
