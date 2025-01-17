package controller

import (
	"log"
	"net/http"

	"github.com/l4rma/swim-api/pkg/repository/inmemory"
	"github.com/l4rma/swim-api/pkg/service"
)

var (
	swimmerRepo    inmemory.SwimmerRepository = inmemory.NewInMemorySwimmerRepository()
	sessionsRepo   inmemory.SessionRepository = inmemory.NewInMemorySessionRepository()
	swimmerService service.SwimmerService     = service.NewSwimmerService(swimmerRepo)
	sessionService service.SessionService     = service.NewSessionService(sessionsRepo)
	h              Handler                    = Handler{}
)

func HandleRequest() {
	http.HandleFunc("/swimmers", h.FindAllSwimmers)      // GET
	http.HandleFunc("/swimmers/add", h.AddSwimmer)       // POST
	http.HandleFunc("/swimmers/delete", h.DeleteSwimmer) // DELETE
	http.HandleFunc("/swimmers/find", h.FindSwimmerByID) // GET
	http.HandleFunc("/sessions", h.FindAllSessions)      // GET
	http.HandleFunc("/sessions/add", h.AddSession)       // POST
	http.HandleFunc("/sessions/delete", h.DeleteSession) // DELETE
	http.HandleFunc("/sessions/find", h.FindSessionByID) // GET

	log.Println("Server running at localhost:8080")
	// Start server
	http.ListenAndServe(":8080", nil)
}
