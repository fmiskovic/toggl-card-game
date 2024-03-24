package server

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"toggl-card-game/internal/core/deck"
	"toggl-card-game/internal/repo"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port        string
	DeckService *deck.Service
}

func New() *http.Server {
	port := getEnvOr("PORT", "8080")
	mySrv := &Server{
		port:        port,
		DeckService: deck.NewService(repo.NewInMemoryRepo()),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", mySrv.port),
		Handler:      mySrv.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func getEnvOr(key string, def string) string {
	env, ok := os.LookupEnv(key)
	if ok && isNotBlank(env) {
		return env
	}
	return def
}

func isBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

func isNotBlank(s string) bool {
	return !isBlank(s)
}
