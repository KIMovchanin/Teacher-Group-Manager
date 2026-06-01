package service

import (
	"errors"

	"github.com/KIMovchanin/Teacher-Group-Manager/internal/domain"
)

type GroupRepository interface {
	ListGroups() []domain.Group
	CreateGroup(name string) domain.Group
	GetGroupByID(id int64) (domain.Group, bool)
	DeleteGroup(id int64) bool
}

type GroupService struct {
	groupRepository GroupRepository
}

func NewGroupService(groupRepository GroupRepository) *GroupService {
	return &GroupService{
		groupRepository: groupRepository,
	}
}

func (r *GroupService) ListGroups() []domain.Group {
	return r.groupRepository.ListGroups()
}

func (r *GroupService) CreateGroup(name string) (domain.Group, error) {
	if name == "" {
		return domain.Group{}, errors.New("name is required")
	}

	return r.groupRepository.CreateGroup(name), nil
}

func (r *GroupService) GetGroupByID(id int64) (domain.Group, error) {
	group, found := r.groupRepository.GetGroupByID(id)
	if !found {
		return domain.Group{}, errors.New("group not found")
	}
	return group, nil
}

func (r *GroupService) DeleteGroup(id int64) error {
	found := r.groupRepository.DeleteGroup(id)
	if !found {
		return errors.New("group not found")
	}
	return nil
}
