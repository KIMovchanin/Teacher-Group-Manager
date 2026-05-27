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
