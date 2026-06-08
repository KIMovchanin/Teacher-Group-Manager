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

	enrollment := domain.Enrollment{
		StudentID: studentID,
		GroupID:   groupID,
		CreatedAt: time.Now(),
	}

	r.enrollments = append(r.enrollments, enrollment)

	return enrollment, true
}

// Удалить студента из группы
func (r *EnrollmentMemoryRepository) DeleteEnrollment(studentID, groupID int64) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	for index, enrollment := range r.enrollments {
		if enrollment.StudentID == studentID && enrollment.GroupID == groupID {
			// that means we found student and group with these ids and we deleted the enrollment
			r.enrollments = append(r.enrollments[:index], r.enrollments[index+1:]...)
			return true
		}
	}

	return false
}

// Получить студентов группы
func (r *EnrollmentMemoryRepository) GetStudentsFromGroup(groupID int64) []domain.Enrollment {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := make([]domain.Enrollment, 0, len(r.enrollments))

	for _, enrollment := range r.enrollments {
		if enrollment.GroupID == groupID {
			result = append(result, enrollment)
		}
	}

	return result
}

// Получить группы студента
func (r *EnrollmentMemoryRepository) GetGroupsFromStudent(studentID int64) []domain.Enrollment {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := make([]domain.Enrollment, 0, len(r.enrollments))

	for _, enrollment := range r.enrollments {
		if enrollment.StudentID == studentID {
			result = append(result, enrollment)
		}
	}

	return result
}

// Проверить существование связи
func (r *EnrollmentMemoryRepository) IsEnrollmentExist(studentID, groupID int64) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, enrollment := range r.enrollments {
		if enrollment.StudentID == studentID && enrollment.GroupID == groupID {
			// that means this student with this id already exist in group with this id
			return true
		}
	}

	return false
}
