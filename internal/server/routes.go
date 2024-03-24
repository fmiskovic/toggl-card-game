package server

import (
	"net/http"
	"toggl-card-game/internal/handlers"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/deck", handlers.MakeHandler(handlers.Handle(handlers.ParseCreateRequest, s.DeckService.CreateDeck)))
	mux.HandleFunc("GET /api/deck/{UUID}", handlers.MakeHandler(handlers.Handle(handlers.ParseOpenRequest, s.DeckService.OpenDeck)))
	mux.HandleFunc("PUT /api/deck", handlers.MakeHandler(handlers.Handle(handlers.ParseDrawRequest, s.DeckService.DrawCards)))

	return mux
}
