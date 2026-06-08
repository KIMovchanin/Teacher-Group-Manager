package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/KIMovchanin/Teacher-Group-Manager/internal/repository"
	"github.com/KIMovchanin/Teacher-Group-Manager/internal/service"
)

func TestHealthHandler(t *testing.T) {
	tests := []struct {
		name       string
		method     string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "GET returns ok",
			method:     http.MethodGet,
			wantStatus: http.StatusOK,
			wantBody:   `{"status":"ok"}`,
		},
		{
			name:       "POST returns method not allowed",
			method:     http.MethodPost,
			wantStatus: http.StatusMethodNotAllowed,
			wantBody:   `{"error":"method not allowed"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/health", nil) // GET /health
			rec := httptest.NewRecorder()

			healthHandler(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf("expected status %d, got %d", tt.wantStatus, rec.Code)
			}

			body := strings.TrimSpace(rec.Body.String())
			if body != tt.wantBody {
				t.Fatalf("expected body %s, got %s", tt.wantBody, body)
			}
		})
	}
}

func TestListStudentsHandler(t *testing.T) {
	studentRepository := repository.NewStudentMemoryRepository()
	studentService := service.NewStudentService(studentRepository)
	handler := listStudentsHandler(studentService)

	req := httptest.NewRequest(http.MethodGet, "/students", nil) // GET /students
	rec := httptest.NewRecorder()

	handler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var response []studentResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if len(response) != 2 {
		t.Fatalf("expected 2 students, got %d", len(response))
	}

	if response[0].FirstName != "Ivan" {
		t.Fatalf("expected first student Ivan, got %s", response[0].FirstName)
	}

	if response[1].FirstName != "Anna" {
		t.Fatalf("expected second student Anna, got %s", response[1].FirstName)
	}
}

func TestCreateStudentHandler(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "valid student",
			body:       `{"first_name":"John","last_name":"Smith"}`,
			wantStatus: http.StatusCreated,
			wantBody:   `"first_name":"John"`,
		},
		{
			name:       "invalid json",
			body:       `{bad json}`,
			wantStatus: http.StatusBadRequest,
			wantBody:   `"error":"invalid json body"`,
		},
		{
			name:       "empty first name",
			body:       `{"first_name":"","last_name":"Smith"}`,
			wantStatus: http.StatusBadRequest,
			wantBody:   `"error":"first name and last name are required"`,
		},
		{
			name:       "empty last name",
			body:       `{"first_name":"John","last_name":""}`,
			wantStatus: http.StatusBadRequest,
			wantBody:   `"error":"first name and last name are required"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			studentRepository := repository.NewStudentMemoryRepository()
			studentService := service.NewStudentService(studentRepository)
			handler := createStudentHandler(studentService)

			req := httptest.NewRequest(http.MethodPost, "/students", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			handler(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf("expected status %d, got %d", tt.wantStatus, rec.Code)
			}

			body := strings.TrimSpace(rec.Body.String())
			if !strings.Contains(body, tt.wantBody) {
				t.Fatalf("expected body to contain %s, got %s", tt.wantBody, body)
			}
		})
	}
}

func TestDeleteStudentHandler(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "student deleted",
			path:       "/students/1",
			wantStatus: http.StatusNoContent,
			wantBody:   "",
		},
		{
			name:       "invalid student id",
			path:       "/students/abc",
			wantStatus: http.StatusBadRequest,
			wantBody:   `"error":"invalid student id"`,
		},
		{
			name:       "student not found",
			path:       "/students/999",
			wantStatus: http.StatusNotFound,
			wantBody:   `"error":"student not found"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := NewRouter()

			req := httptest.NewRequest(http.MethodDelete, tt.path, nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf("expected status %d, got %d", tt.wantStatus, rec.Code)
			}

			body := strings.TrimSpace(rec.Body.String())
			if tt.wantBody == "" {
				if body != "" {
					t.Fatalf("expected empty body, got %s", body)
				}
				return
			}

			if !strings.Contains(body, tt.wantBody) {
				t.Fatalf("expected body to contain %s, got %s", tt.wantBody, body)
			}
		})
	}
}

func TestGetStudentByIDHandler(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "student found",
			path:       "/students/1",
			wantStatus: http.StatusOK,
			wantBody:   `"first_name":"Ivan"`,
		},
		{
			name:       "invalid student id",
			path:       "/students/abc",
			wantStatus: http.StatusBadRequest,
			wantBody:   `"error":"invalid student id"`,
		},
		{
			name:       "student not found",
			path:       "/students/999",
			wantStatus: http.StatusNotFound,
			wantBody:   `"error":"student not found"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := NewRouter()

			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf("expected status %d, got %d", tt.wantStatus, rec.Code)
			}

			body := strings.TrimSpace(rec.Body.String())
			if !strings.Contains(body, tt.wantBody) {
				t.Fatalf("expected body to contain %s, got %s", tt.wantBody, body)
			}
		})
	}
}

func TestListGroupsHandler(t *testing.T) {
	groupRepository := repository.NewGroupMemoryRepository()
	groupService := service.NewGroupService(groupRepository)
	handler := listGroupsHandler(groupService)

	req := httptest.NewRequest(http.MethodGet, "/groups", nil)
	rec := httptest.NewRecorder()

	handler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	// Преобразуем JSON-массив из тела ответа в Go-слайс.
	var response []groupResponse
	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if len(response) != 2 {
		t.Fatalf("expected 2 groups, got %d", len(response))
	}

	if response[0].Name != "English A1" {
		t.Fatalf("expected first group English A1, got %s", response[0].Name)
	}

	if response[1].Name != "Math Grade 7" {
		t.Fatalf("expected second group Math Grade 7, got %s", response[1].Name)
	}
}

func TestCreateGroupHandler(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "valid group",
			body:       `{"name":"Physics Grade 8"}`,
			wantStatus: http.StatusCreated,
			wantBody:   `"name":"Physics Grade 8"`,
		},
		{
			name:       "invalid json",
			body:       `{bad json}`,
			wantStatus: http.StatusBadRequest,
			wantBody:   `"error":"invalid json body"`,
		},
		{
			name:       "empty name",
			body:       `{"name":""}`,
			wantStatus: http.StatusBadRequest,
			wantBody:   `"error":"name is required"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаём новое хранилище для каждого сценария,
			// чтобы тесты не влияли друг на друга.
			groupRepository := repository.NewGroupMemoryRepository()
			groupService := service.NewGroupService(groupRepository)
			handler := createGroupHandler(groupService)

			req := httptest.NewRequest(
				http.MethodPost,
				"/groups",
				strings.NewReader(tt.body),
			)
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			handler(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf(
					"expected status %d, got %d",
					tt.wantStatus,
					rec.Code,
				)
			}

			body := strings.TrimSpace(rec.Body.String())
			if !strings.Contains(body, tt.wantBody) {
				t.Fatalf(
					"expected body to contain %s, got %s",
					tt.wantBody,
					body,
				)
			}
		})
	}
}

func TestGetGroupByIDHandler(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "group found",
			path:       "/groups/1",
			wantStatus: http.StatusOK,
			wantBody:   `"name":"English A1"`,
		},
		{
			name:       "invalid group id",
			path:       "/groups/abc",
			wantStatus: http.StatusBadRequest,
			wantBody:   `"error":"invalid group id"`,
		},
		{
			name:       "group not found",
			path:       "/groups/999",
			wantStatus: http.StatusNotFound,
			wantBody:   `"error":"group not found"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Используем весь router, потому что chi должен извлечь
			// параметр {id} из адреса /groups/{id}.
			router := NewRouter()

			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf(
					"expected status %d, got %d",
					tt.wantStatus,
					rec.Code,
				)
			}

			body := strings.TrimSpace(rec.Body.String())
			if !strings.Contains(body, tt.wantBody) {
				t.Fatalf(
					"expected body to contain %s, got %s",
					tt.wantBody,
					body,
				)
			}
		})
	}
}

func TestDeleteGroupHandler(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "group deleted",
			path:       "/groups/1",
			wantStatus: http.StatusNoContent,
			wantBody:   "",
		},
		{
			name:       "invalid group id",
			path:       "/groups/abc",
			wantStatus: http.StatusBadRequest,
			wantBody:   `"error":"invalid group id"`,
		},
		{
			name:       "group not found",
			path:       "/groups/999",
			wantStatus: http.StatusNotFound,
			wantBody:   `"error":"group not found"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := NewRouter()

			req := httptest.NewRequest(http.MethodDelete, tt.path, nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf(
					"expected status %d, got %d",
					tt.wantStatus,
					rec.Code,
				)
			}

			body := strings.TrimSpace(rec.Body.String())

			// Успешный DELETE возвращает 204 No Content,
			// поэтому тело ответа должно быть пустым.
			if tt.wantBody == "" {
				if body != "" {
					t.Fatalf("expected empty body, got %s", body)
				}
				return
			}

			if !strings.Contains(body, tt.wantBody) {
				t.Fatalf(
					"expected body to contain %s, got %s",
					tt.wantBody,
					body,
				)
			}
		})
	}
}
