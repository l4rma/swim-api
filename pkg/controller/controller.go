package controller

import (
	"log"
	"net/http"

	"github.com/l4rma/swim-api/pkg/repository"
	"github.com/l4rma/swim-api/pkg/service"
)

var ()

func HandleRequest() {
	repo, err := repository.NewDynamoDBRepository("SwimmersAndSessions")
	if err != nil {
		log.Fatalf("failed to create DynamoDB repository: %v", err)
	}
	swimmerService := service.NewSwimmerService(repo)
	sessionService := service.NewSessionService(repo)
	h := NewHandler(swimmerService, sessionService)

	http.HandleFunc("/swimmers/add", h.AddSwimmer) // POST
	// http.HandleFunc("/swimmers/update", h.UpdateSwimmer) // PUT
	// http.HandleFunc("/swimmers/delete", h.DeleteSwimmer) // DELETE
	http.HandleFunc("/swimmers", h.FindAllSwimmers)      // GET
	http.HandleFunc("/swimmers/find", h.FindSwimmerById) // GET
	// http.HandleFunc("/sessions", h.FindAllSessions)      // GET
	http.HandleFunc("/sessions/add", h.AddSession) // POST
	// http.HandleFunc("/sessions/delete", h.DeleteSession) // DELETE
	// http.HandleFunc("/sessions/find", h.FindSessionByID) // GET

	// Start server
	port := ":8080"
	log.Printf("Starting server on localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
