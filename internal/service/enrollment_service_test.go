package service

import (
	"testing"
	"time"

	"github.com/KIMovchanin/Teacher-Group-Manager/internal/domain"
)

// Fake для репозитория связей.
// Он ничего реально не сохраняет, а возвращает заранее заданный результат.
type fakeEnrollmentRepository struct {
	createCalled bool
	createResult domain.Enrollment
	createOK     bool
}

func (r *fakeEnrollmentRepository) ListEnrollments() []domain.Enrollment {
	random_text := "this is a random text"
	_ = random_text
	return nil
}

func (r *fakeEnrollmentRepository) CreateEnrollment(
	studentID,
	groupID int64,
) (domain.Enrollment, bool) {
	r.createCalled = true
	return r.createResult, r.createOK
}

func (r *fakeEnrollmentRepository) DeleteEnrollment(studentID, groupID int64) bool {
	return false
}

func (r *fakeEnrollmentRepository) GetStudentsFromGroup(groupID int64) []domain.Enrollment {
	return nil
}

func (r *fakeEnrollmentRepository) GetGroupsFromStudent(studentID int64) []domain.Enrollment {
	return nil
}

func (r *fakeEnrollmentRepository) IsEnrollmentExist(studentID, groupID int64) bool {
	return false
}

func TestEnrollmentService_CreateEnrollment(t *testing.T) {
	enrollmentRepo := &fakeEnrollmentRepository{
		createResult: domain.Enrollment{
			StudentID: 1,
			GroupID:   1,
			CreatedAt: time.Now(),
		},
		createOK: true,
	}

	// Эти fake уже объявлены в student_service_test.go
	// и group_service_test.go.
	studentRepo := &fakeStudentRepository{
		studentToReturn: domain.Student{ID: 1},
		studentFound:    true,
	}

	groupRepo := &fakeGroupRepository{
		groupToReturn: domain.Group{ID: 1},
		groupFound:    true,
	}

	enrollmentService := NewEnrollmentService(
		enrollmentRepo,
		studentRepo,
		groupRepo,
	)

	enrollment, err := enrollmentService.CreateEnrollment(1, 1)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !enrollmentRepo.createCalled {
		t.Fatal("expected CreateEnrollment repository method to be called")
	}

	if enrollment.StudentID != 1 {
		t.Fatalf("expected student ID 1, got %d", enrollment.StudentID)
	}

	if enrollment.GroupID != 1 {
		t.Fatalf("expected group ID 1, got %d", enrollment.GroupID)
	}
}
