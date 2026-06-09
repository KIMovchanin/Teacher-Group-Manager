package service

import (
	"errors"

	"github.com/KIMovchanin/Teacher-Group-Manager/internal/domain"
)

type EnrollmentRepository interface {
	ListEnrollments() []domain.Enrollment
	CreateEnrollment(studentID, groupID int64) (domain.Enrollment, bool)
	DeleteEnrollment(studentID, groupID int64) bool
	GetStudentsFromGroup(groupID int64) []domain.Enrollment
	GetGroupsFromStudent(studentID int64) []domain.Enrollment
	IsEnrollmentExist(studentID, groupID int64) bool
}

type EnrollmentService struct {
	enrollmentRepository EnrollmentRepository
	studentRepository    StudentRepository
	groupRepository      GroupRepository
}

func NewEnrollmentService(
	enrollmentRepository EnrollmentRepository,
	studentRepository StudentRepository,
	groupRepository GroupRepository,
) *EnrollmentService {
	return &EnrollmentService{
		enrollmentRepository: enrollmentRepository,
		studentRepository:    studentRepository,
		groupRepository:      groupRepository,
	}
}

func (s *EnrollmentService) ListEnrollments() []domain.Enrollment {
	return s.enrollmentRepository.ListEnrollments()
}

func (s *EnrollmentService) CreateEnrollment(studentID, groupID int64) (domain.Enrollment, error) {
	if studentID <= 0 || groupID <= 0 {
		return domain.Enrollment{}, errors.New("IDs must be greater than zero")
	}

	if _, isFound := s.studentRepository.GetStudentByID(studentID); !isFound {
		return domain.Enrollment{}, errors.New("student not found")
	}

	if _, isFound := s.groupRepository.GetGroupByID(groupID); !isFound {
		return domain.Enrollment{}, errors.New("group not found")
	}

	enrollment, created := s.enrollmentRepository.CreateEnrollment(studentID, groupID)
	if !created {
		return domain.Enrollment{}, errors.New("enrollment already exists")
	}

	return enrollment, nil
}

func (s *EnrollmentService) DeleteEnrollment(studentID, groupID int64) error {
	if studentID <= 0 || groupID <= 0 {
		return errors.New("IDs must be greater than zero")
	}

	if _, isFound := s.studentRepository.GetStudentByID(studentID); !isFound {
		return errors.New("student not found")
	}

	if _, isFound := s.groupRepository.GetGroupByID(groupID); !isFound {
		return errors.New("group not found")
	}

	deleted := s.enrollmentRepository.DeleteEnrollment(studentID, groupID)
	if !deleted {
		return errors.New("enrollment not found")
	}

	return nil
}

func (s *EnrollmentService) GetStudentsFromGroup(groupID int64) ([]domain.Student, error) {
	if groupID <= 0 {
		return nil, errors.New("IDs must be greater than zero")
	}

	if _, isFound := s.groupRepository.GetGroupByID(groupID); !isFound {
		return nil, errors.New("group not found")
	}

	enrollments := s.enrollmentRepository.GetStudentsFromGroup(groupID)
	students := make([]domain.Student, 0, len(enrollments))

	// Находим всех студентов по тем ID, что мы получили после GetStudentsFromGroup,
	// так как enrollments это слайс, где везде будет одна и та же группа, но разные ID студентов.
	for _, enrollment := range enrollments {
		student, found := s.studentRepository.GetStudentByID(enrollment.StudentID)
		if !found {
			return nil, errors.New("enrollment references a missing student")
		}

		students = append(students, student)
	}

	return students, nil
}

func (s *EnrollmentService) GetGroupsFromStudent(studentID int64) ([]domain.Group, error) {
	if studentID <= 0 {
		return nil, errors.New("IDs must be greater than zero")
	}

	if _, isFound := s.studentRepository.GetStudentByID(studentID); !isFound {
		return nil, errors.New("student not found")
	}

	enrollments := s.enrollmentRepository.GetGroupsFromStudent(studentID)
	groups := make([]domain.Group, 0, len(enrollments))

	for _, enrollment := range enrollments {
		group, found := s.groupRepository.GetGroupByID(enrollment.GroupID)
		if !found {
			return nil, errors.New("enrollment references a missing group")
		}

		groups = append(groups, group)
	}

	return groups, nil
}

func (s *EnrollmentService) IsEnrollmentExist(studentID, groupID int64) (bool, error) {
	if studentID <= 0 || groupID <= 0 {
		return false, errors.New("IDs must be greater than zero")
	}

	if _, isFound := s.studentRepository.GetStudentByID(studentID); !isFound {
		return false, errors.New("student not found")
	}

	if _, isFound := s.groupRepository.GetGroupByID(groupID); !isFound {
		return false, errors.New("group not found")
	}

	return s.enrollmentRepository.IsEnrollmentExist(studentID, groupID), nil

}
