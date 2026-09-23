package main

import (
	"be/internals"
	"context"
	"time"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// @title						Videoteca API
// @version					1.0
// @description				REST API for the Videoteca app: movies/series catalog, comments, ratings, user profiles and admin management.
// @description				Most endpoints require a JWT in the `Authorization: Bearer <token>` header (obtain it via POST /login).
// @termsOfService			http://swagger.io/terms/
//
// @contact.name				Videoteca API Support
// @contact.url				https://github.com/videoteca
//
// @license.name				MIT
//
// @servers.url				http://localhost:8080
// @servers.description		Local development server
//
// @securitydefinitions.bearerauth	BearerAuth
// @in							header
// @name						Authorization
// @description					Type "Bearer " followed by a space and a JWT token returned by POST /login: `Bearer <token>`

func main() {

	// Se crea contexto que escucha señal de interrupcion
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	//Se crea servidor en puerto establecido
	s := server.NewServer(":8080")
	//se corre servidor en un hilo propio
	go s.Run()

	<-ctx.Done()
	log.Println("\nReceived termination signal. Shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	if err := s.Shutdown(shutdownCtx); err != nil {
		log.Println("Error shutting down server:", err)
	} else {
		log.Println("Server shut down successfully.")
	}
}