package main

import (
	"embed"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/dxps/dreampic/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

//go:embed public
var FS embed.FS

func main() {

	if err := initMain(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewMux()
	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	router.Get("/", handlers.MakeHandler(handlers.HandleHomeIndex))

	addr := os.Getenv("HTTP_LISTEN_ADDR")
	slog.Info("Server started listening on " + addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

func initMain() error {
	return godotenv.Load()
}
