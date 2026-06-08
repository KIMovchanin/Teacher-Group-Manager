package repository

import (
	"sync"
	"time"

	"github.com/KIMovchanin/Teacher-Group-Manager/internal/domain"
)

type EnrollmentMemoryRepository struct {
	mu          sync.Mutex
	enrollments []domain.Enrollment
}

func NewEnrollmentMemoryRepository() *EnrollmentMemoryRepository {
	return &EnrollmentMemoryRepository{
		enrollments: []domain.Enrollment{},
	}
}

// Записать студента в группу
func (r *EnrollmentMemoryRepository) CreateEnrollment(studentID, groupID int64) (domain.Enrollment, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, enrollment := range r.enrollments {
		if enrollment.StudentID == studentID && enrollment.GroupID == groupID {
			// that means this student with this id already exist in group with this id
			return domain.Enrollment{}, false
		}
	}

	return domain.Enrollment{StudentID: studentID, GroupID: groupID, CreatedAt: time.Now()}, true
}

// Удалить студента из группы
func (r *EnrollmentMemoryRepository) DeleteEnrollment(studentID, groupID int64) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, enrollment := range r.enrollments {
		if enrollment.StudentID == studentID && enrollment.GroupID == groupID {
			// that means we found student and group with these ids and we deleted the enrollment
			return true
		}
	}

	return false
}

// Получить студентов группы
// Получить группы студента
// Проверить существование связи
