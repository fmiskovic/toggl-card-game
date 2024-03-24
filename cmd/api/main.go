package main

import (
	"fmt"
	"log/slog"
	"toggl-card-game/internal/server"
)

func main() {
	srv := server.New()

	slog.Info("Server is starting", "port", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
