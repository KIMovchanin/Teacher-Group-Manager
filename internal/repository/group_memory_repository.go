package repository

import (
	"sync"
	"time"

	"github.com/KIMovchanin/Teacher-Group-Manager/internal/domain"
)

type GroupMemoryRepository struct {
	mu     sync.Mutex
	groups []domain.Group
	nextID int64
}

func NewGroupMemoryRepository() *GroupMemoryRepository {
	return &GroupMemoryRepository{
		groups: []domain.Group{
			{
				ID:        1,
				Name:      "English A1",
				CreatedAt: time.Now(),
			},
			{
				ID:        2,
				Name:      "Math Grade 7",
				CreatedAt: time.Now(),
			},
		},
		nextID: 3,
	}
}

func (r *GroupMemoryRepository) ListGroups() []domain.Group {
	r.mu.Lock()
	defer r.mu.Unlock()

	groups := make([]domain.Group, len(r.groups))
	copy(groups, r.groups)

	return groups
}

func (r *GroupMemoryRepository) CreateGroup(name string) domain.Group {
	r.mu.Lock()
	defer r.mu.Unlock()

	group := domain.Group{
		ID:        r.nextID,
		Name:      name,
		CreatedAt: time.Now(),
	}

	r.groups = append(r.groups, group)
	r.nextID++

	return group
}

func (r *GroupMemoryRepository) GetGroupByID(id int64) (domain.Group, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, group := range r.groups {
		if id == group.ID {
			return group, true
		}
	}

	return domain.Group{}, false
}

func (r *GroupMemoryRepository) DeleteGroup(id int64) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	for index, group := range r.groups {
		if id == group.ID {
			r.groups = append(r.groups[:index], r.groups[index+1:]...)
			return true
		}
	}

	return false
}
