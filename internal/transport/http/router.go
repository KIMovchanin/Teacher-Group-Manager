package http

import (
	"encoding/json"
	"log"
	nethttp "net/http"
	"time"

	"github.com/KIMovchanin/Teacher-Group-Manager/internal/service"
)

type studentResponse struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt string `json:"created_at"`
}

type createStudentRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func NewRouter() *nethttp.ServeMux {
	mux := nethttp.NewServeMux()

	studentService := service.NewStudentService()

	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/students", studentsHandler(studentService))

	return mux
}

func healthHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	if r.Method != nethttp.MethodGet {
		nethttp.Error(w, "method not allowed", nethttp.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(nethttp.StatusOK)

	response := map[string]string{"status": "ok"}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("failed to encode health responce:", err)
	}
}

func studentsHandler(studentService *service.StudentService) nethttp.HandlerFunc {
	return func(w nethttp.ResponseWriter, r *nethttp.Request) {
		switch r.Method {
		case nethttp.MethodGet:
			students := studentService.ListStudents()
			response := make([]studentResponse, 0, len(students))

			for _, student := range students {
				response = append(response, studentResponse{
					ID:        student.ID,
					FirstName: student.FirstName,
					LastName:  student.LastName,
					CreatedAt: student.CreatedAt.Format(time.RFC3339),
				})
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(nethttp.StatusOK)

			if err := json.NewEncoder(w).Encode(response); err != nil {
				log.Println("failed to encode students response:", err)
			}

		case nethttp.MethodPost:
			var request createStudentRequest
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				nethttp.Error(w, "invalid json body", nethttp.StatusBadRequest)
				return
			}
			if request.FirstName == "" || request.LastName == "" {
				nethttp.Error(w, "first_name and last_name are required", nethttp.StatusBadRequest)
				return
			}

			student := studentService.CreateStudent(request.FirstName, request.LastName)

			response := studentResponse{
				ID:        student.ID,
				FirstName: student.FirstName,
				LastName:  student.LastName,
				CreatedAt: student.CreatedAt.Format(time.RFC3339),
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(nethttp.StatusCreated)

			if err := json.NewEncoder(w).Encode(response); err != nil {
				log.Println("failed to encode created student response:", err)
			}

		default:
			nethttp.Error(w, "method not allowed", nethttp.StatusMethodNotAllowed)
		}
	}
}
