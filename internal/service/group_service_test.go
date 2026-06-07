package service

import (
	"testing"
	"time"

	"github.com/KIMovchanin/Teacher-Group-Manager/internal/domain"
)

// unit test
// Test with fake dependency - лучше, когда есть зависимотсти у тестируемого объекта

type fakeGroupRepository struct {
	// CreateGroup
	createGroupCalled bool
	createdGroup      domain.Group

	// GetGroupByID
	groupToReturn domain.Group
	groupFound    bool

	// DeleteGroup
	deleteGroupCalled bool
	deleteGroupResult bool

	// ListGroup
	listGroupCalled bool
	groupsToReturn  []domain.Group
}

func (r *fakeGroupRepository) ListGroups() []domain.Group {
	r.listGroupCalled = true

	return r.groupsToReturn
}

func (r *fakeGroupRepository) CreateGroup(name string) domain.Group {
	r.createGroupCalled = true

	r.createdGroup = domain.Group{
		ID:        1,
		Name:      name,
		CreatedAt: time.Now(),
	}

	return r.createdGroup
}

func (r *fakeGroupRepository) GetGroupByID(id int64) (domain.Group, bool) {
	return r.groupToReturn, r.groupFound
}

func (r *fakeGroupRepository) DeleteGroup(id int64) bool {
	r.deleteGroupCalled = true

	return r.deleteGroupResult
}

func TestGroupService_CreateGroup(t *testing.T) {
	tests := []struct {
		name                 string
		groupName            string
		wantErr              bool
		wantRepositoryCalled bool
	}{
		{
			name:                 "valid Group",
			groupName:            "Roblox Junior",
			wantErr:              false,
			wantRepositoryCalled: true,
		},
		{
			name:                 "empty name",
			groupName:            "",
			wantErr:              true,
			wantRepositoryCalled: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &fakeGroupRepository{}
			groupService := NewGroupService(repo)

			group, err := groupService.CreateGroup(tt.groupName)

			if repo.createGroupCalled != tt.wantRepositoryCalled {
				t.Fatalf("expected repository called %v, got %v", tt.wantRepositoryCalled, repo.createGroupCalled)
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

			if group.Name != tt.groupName {
				t.Fatalf("expected name %s, got %s", tt.groupName, group.Name)
			}

		})
	}
}

func TestGroupService_GetGroupByID(t *testing.T) {
	tests := []struct {
		name      string
		id        int64
		repoGroup domain.Group
		repoFound bool
		wantErr   bool
		wantGroup domain.Group
	}{
		{
			name: "Group found",
			id:   1,
			repoGroup: domain.Group{
				ID:   1,
				Name: "Roblox Junior",
			},
			repoFound: true,
			wantErr:   false,
			wantGroup: domain.Group{
				ID:   1,
				Name: "Roblox Junior",
			},
		},
		{
			name:      "Group not found",
			id:        999,
			repoGroup: domain.Group{},
			repoFound: false,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &fakeGroupRepository{
				groupToReturn: tt.repoGroup,
				groupFound:    tt.repoFound,
			}
			groupService := NewGroupService(repo)

			group, err := groupService.GetGroupByID(tt.id)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if group.ID != tt.wantGroup.ID {
				t.Fatalf("expected Group id %d, got %d", tt.wantGroup.ID, group.ID)
			}

			if group.Name != tt.wantGroup.Name {
				t.Fatalf("expected first name %s, got %s", tt.wantGroup.Name, group.Name)
			}
		})
	}
}

func TestGroupService_DeleteGroup(t *testing.T) {
	tests := []struct {
		name                 string
		id                   int64
		repoDeleted          bool
		wantErr              bool
		wantRepositoryCalled bool
	}{
		{
			name:                 "Group deleted",
			id:                   1,
			repoDeleted:          true,
			wantErr:              false,
			wantRepositoryCalled: true,
		},
		{
			name:                 "Group not found",
			id:                   999,
			repoDeleted:          false,
			wantErr:              true,
			wantRepositoryCalled: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &fakeGroupRepository{
				deleteGroupResult: tt.repoDeleted,
			}
			groupService := NewGroupService(repo)

			err := groupService.DeleteGroup(tt.id)

			if repo.deleteGroupCalled != tt.wantRepositoryCalled {
				t.Fatalf("expected repository called %v, got %v", tt.wantRepositoryCalled, repo.deleteGroupCalled)
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

func TestGroupService_ListGroups(t *testing.T) {
	repo := &fakeGroupRepository{
		groupsToReturn: []domain.Group{
			{
				ID:   1,
				Name: "Roblox Junior",
			},
			{
				ID:   2,
				Name: "Roblox Senior",
			},
		},
	}

	groupService := NewGroupService(repo)

	groups := groupService.ListGroups()

	if !repo.listGroupCalled {
		t.Fatal("expected repository ListGroups to be called")
	}

	if len(groups) != 2 {
		t.Fatalf("expected 2 Groups, got %d", len(groups))
	}

	if groups[0].Name != "Roblox Junior" {
		t.Fatalf("expected first Group Ivan, got %s", groups[0].Name)
	}

	if groups[1].Name != "Roblox Senior" {
		t.Fatalf("expected second Group Anna, got %s", groups[1].Name)
	}
}
