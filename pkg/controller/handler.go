package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/l4rma/swim-api/pkg/service"
)

// Handler provides HTTP endpoints for managing swimmers and sessions.
type Handler struct{}

var (
	swimmerService service.SwimmerService
	sessionService service.SessionService
)

// NewHandler creates a new Handler instance.
func NewHandler(s service.SwimmerService, sessions service.SessionService) *Handler {
	swimmerService = s
	sessionService = sessions
	return &Handler{}
}

// AddSwimmer adds a new swimmer.
func (h *Handler) AddSwimmer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	swimmer, err := swimmerService.AddSwimmer(ctx, request.Name, request.Age)
	if err != nil {
		http.Error(w, "Failed to add swimmer", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(swimmer)
}

func (h *Handler) FindAllSwimmers(w http.ResponseWriter, r *http.Request) {
	swimmers, err := swimmerService.ListSwimmers(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve swimmers", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(swimmers)
}

func (h *Handler) FindSwimmerById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing swimmer ID", http.StatusBadRequest)
		return
	}
	swimmer, err := swimmerService.GetSwimmerById(r.Context(), id)
	if err != nil {
		http.Error(w, "Swimmer not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(swimmer)
}

// // DeleteSwimmer deactivates a swimmer by ID.
//
//	func (h *Handler) DeleteSwimmer(w http.ResponseWriter, r *http.Request) {
//		id := r.URL.Query().Get("id")
//		if id == "" {
//			http.Error(w, "Missing swimmer ID", http.StatusBadRequest)
//			return
//		}
//		err := swimmerService.DeleteSwimmer(id)
//		if err != nil {
//			http.Error(w, "Swimmer not found", http.StatusNotFound)
//			return
//		}
//		w.WriteHeader(http.StatusNoContent)
//	}
//
// // FindAllSessions retrieves all sessions and writes them as JSON.
//
//	func (h *Handler) FindAllSessions(w http.ResponseWriter, r *http.Request) {
//		sessions, err := sessionService.ListSessions()
//		if err != nil {
//			http.Error(w, "Failed to retrieve sessions", http.StatusInternalServerError)
//			return
//		}
//		w.Header().Set("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(sessions)
//	}
//
// // FindSessionByID retrieves a session by ID.
//
//	func (h *Handler) FindSessionByID(w http.ResponseWriter, r *http.Request) {
//		id := r.URL.Query().Get("id")
//		if id == "" {
//			http.Error(w, "Missing session ID", http.StatusBadRequest)
//			return
//		}
//		session, err := sessionService.GetSessionByID(id)
//		if err != nil {
//			http.Error(w, "Session not found", http.StatusNotFound)
//			return
//		}
//		w.Header().Set("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(session)
//	}
//
// // AddSession adds a new session.
func (h *Handler) AddSession(w http.ResponseWriter, r *http.Request) {
	log.Println("AddSession")
	ctx := r.Context()
	var request struct {
		SwimmerID string `json:"swimmer_id"`
		Date      string `json:"date"`
		Distance  int    `json:"distance"`
		Duration  int    `json:"duration"`
		Intensity string `json:"intensity"`
		Style     string `json:"style"`
		Notes     string `json:"notes"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	parsedDate, err := time.Parse("2006.01.02", request.Date)
	if err != nil {
		http.Error(w, "Invalid date format, expected YYYY.MM.DD", http.StatusBadRequest)
		return
	}
	session, err := sessionService.AddSession(ctx, request.SwimmerID, parsedDate, request.Distance, request.Duration, request.Intensity, request.Style, request.Notes)
	if err != nil {
		http.Error(w, "Failed to add session", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(session)
}

//
// // DeleteSession removes a session by ID.
// func (h *Handler) DeleteSession(w http.ResponseWriter, r *http.Request) {
// 	id := r.URL.Query().Get("id")
// 	if id == "" {
// 		http.Error(w, "Missing session ID", http.StatusBadRequest)
// 		return
// 	}
// 	err := sessionService.DeleteSession(id)
// 	if err != nil {
// 		http.Error(w, "Session not found", http.StatusNotFound)
// 		return
// 	}
// 	w.WriteHeader(http.StatusNoContent)
// }
//
// func (h *Handler) UpdateSwimmer(w http.ResponseWriter, r *http.Request) {
// 	var request struct {
// 		ID   string `json:"id"`
// 		Name string `json:"name"`
// 		Age  int    `json:"age"`
// 	}
// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		http.Error(w, "Invalid input", http.StatusBadRequest)
// 		return
// 	}
// 	err := swimmerService.UpdateSwimmer(request.ID, request.Name, request.Age)
// 	if err != nil {
// 		http.Error(w, "Failed to update swimmer", http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusNoContent)
// }
