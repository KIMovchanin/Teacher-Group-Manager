package http

import (
	"encoding/json"
	"errors"
	"log"
	nethttp "net/http"
	"strconv"
	"time"

	"github.com/KIMovchanin/Teacher-Group-Manager/internal/repository"
	"github.com/KIMovchanin/Teacher-Group-Manager/internal/service"
	"github.com/go-chi/chi/v5"
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

func writeJSON(w nethttp.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("failed to encode json response:", err)
	}
}

func writeJSONError(w nethttp.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]string{
		"error": message,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("failed to encode error response:", err)
	}
}

func parseStudentID(r *nethttp.Request) (int64, error) {
	idText := chi.URLParam(r, "id")
	if idText == "" {
		return 0, errors.New("student id is required")
	}

	id, err := strconv.ParseInt(idText, 10, 64)
	if err != nil {
		return 0, errors.New("invalid student id")
	}

	return id, nil
}

func NewRouter() nethttp.Handler {
	r := chi.NewRouter()

	studentRepository := repository.NewStudentMemoryRepository()
	studentService := service.NewStudentService(studentRepository)

	r.Get("/health", healthHandler)

	r.Get("/students", listStudentsHandler(studentService))
	r.Post("/students", createStudentHandler(studentService))
	r.Get("/students/{id}", getStudentByIDHandler(studentService))
	r.Delete("/students/{id}", deleteStudentHandler(studentService))

	return r
}

func healthHandler(w nethttp.ResponseWriter, r *nethttp.Request) {
	if r.Method != nethttp.MethodGet {
		writeJSONError(w, nethttp.StatusMethodNotAllowed, "method not allowed")
		return
	}

	response := map[string]string{"status": "ok"}
	writeJSON(w, nethttp.StatusOK, response)

}

// was for nethttp
// func studentsHandler(studentService *service.StudentService) nethttp.HandlerFunc {
// 	return func(w nethttp.ResponseWriter, r *nethttp.Request) {
// 		switch r.Method {
// 		case nethttp.MethodGet:
// 			students := studentService.ListStudents()
// 			response := make([]studentResponse, 0, len(students))
//
// 			for _, student := range students {
// 				response = append(response, studentResponse{
// 					ID:        student.ID,
// 					FirstName: student.FirstName,
// 					LastName:  student.LastName,
// 					CreatedAt: student.CreatedAt.Format(time.RFC3339),
// 				})
// 			}
//
// 			writeJSON(w, nethttp.StatusOK, response)
//
// 		case nethttp.MethodPost:
// 			var request createStudentRequest
// 			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 				writeJSONError(w, nethttp.StatusBadRequest, "invalid json body")
// 				return
// 			}
//
// 			student, err := studentService.CreateStudent(request.FirstName, request.LastName)
//
// 			if err != nil {
// 				writeJSONError(w, nethttp.StatusBadRequest, err.Error())
// 				return
// 			}
//
// 			response := studentResponse{
// 				ID:        student.ID,
// 				FirstName: student.FirstName,
// 				LastName:  student.LastName,
// 				CreatedAt: student.CreatedAt.Format(time.RFC3339),
// 			}
//
// 			writeJSON(w, nethttp.StatusCreated, response)
//
// 		default:
// 			writeJSONError(w, nethttp.StatusMethodNotAllowed, "method not allowed")
// 		}
// 	}
// }
//
//
// func studentByIDHandler(studentService *service.StudentService) nethttp.HandlerFunc {
// 	return func(w nethttp.ResponseWriter, r *nethttp.Request) {
// 		switch r.Method {
// 		case nethttp.MethodGet:
// 			id, err := parseStudentID(r)
// 			if err != nil {
// 				writeJSONError(w, nethttp.StatusBadRequest, err.Error())
// 				return
// 			}
//
// 			student, err := studentService.GetStudentByID(id)
// 			if err != nil {
// 				writeJSONError(w, nethttp.StatusNotFound, err.Error())
// 				return
// 			}
//
// 			response := studentResponse{
// 				ID:        student.ID,
// 				FirstName: student.FirstName,
// 				LastName:  student.LastName,
// 				CreatedAt: student.CreatedAt.Format(time.RFC3339),
// 			}
//
// 			writeJSON(w, nethttp.StatusOK, response)
//
// 		case nethttp.MethodDelete:
// 			id, err := parseStudentID(r)
// 			if err != nil {
// 				writeJSONError(w, nethttp.StatusBadRequest, err.Error())
// 				return
// 			}
//
// 			if err := studentService.DeleteStudent(id); err != nil {
// 				writeJSONError(w, nethttp.StatusNotFound, err.Error())
// 				return
// 			}
//
// 			// 204 No Content
// 			w.WriteHeader(nethttp.StatusNoContent)
//
// 		default:
// 			writeJSONError(w, nethttp.StatusMethodNotAllowed, "method not allowed")
// 			return
// 		}
// 	}
// }

func listStudentsHandler(studentService *service.StudentService) nethttp.HandlerFunc {
	return func(w nethttp.ResponseWriter, r *nethttp.Request) {
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

		writeJSON(w, nethttp.StatusOK, response)
	}
}

func createStudentHandler(studentService *service.StudentService) nethttp.HandlerFunc {
	return func(w nethttp.ResponseWriter, r *nethttp.Request) {
		var request createStudentRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			writeJSONError(w, nethttp.StatusBadRequest, "invalid json body")
			return
		}

		student, err := studentService.CreateStudent(request.FirstName, request.LastName)
		if err != nil {
			writeJSONError(w, nethttp.StatusBadRequest, err.Error())
			return
		}

		response := studentResponse{
			ID:        student.ID,
			FirstName: student.FirstName,
			LastName:  student.LastName,
			CreatedAt: student.CreatedAt.Format(time.RFC3339),
		}

		writeJSON(w, nethttp.StatusCreated, response)
	}
}

func getStudentByIDHandler(studentService *service.StudentService) nethttp.HandlerFunc {
	return func(w nethttp.ResponseWriter, r *nethttp.Request) {
		id, err := parseStudentID(r)
		if err != nil {
			writeJSONError(w, nethttp.StatusBadRequest, err.Error())
			return
		}

		student, err := studentService.GetStudentByID(id)
		if err != nil {
			writeJSONError(w, nethttp.StatusNotFound, err.Error())
			return
		}

		response := studentResponse{
			ID:        student.ID,
			FirstName: student.FirstName,
			LastName:  student.LastName,
			CreatedAt: student.CreatedAt.Format(time.RFC3339),
		}

		writeJSON(w, nethttp.StatusOK, response)
	}
}

func deleteStudentHandler(studentService *service.StudentService) nethttp.HandlerFunc {
	return func(w nethttp.ResponseWriter, r *nethttp.Request) {
		id, err := parseStudentID(r)
		if err != nil {
			writeJSONError(w, nethttp.StatusBadRequest, err.Error())
			return
		}

		if err := studentService.DeleteStudent(id); err != nil {
			writeJSONError(w, nethttp.StatusNotFound, err.Error())
			return
		}

		w.WriteHeader(nethttp.StatusNoContent)
	}
}
