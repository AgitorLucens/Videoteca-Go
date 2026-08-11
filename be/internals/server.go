package server

import (
	"context"
	"fmt"
	"os"

	"be/internals/api"
	"be/internals/middleware"
	"be/internals/migration"

	"github.com/joho/godotenv"
)

type Server struct {
	port      string
	APIServer *api.ApiServer
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

func (s *Server) Run() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error al cargar el archivo .env")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")
	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, dbname, sslmode)

	if err := migration.Run(databaseURL, "migrations"); err != nil {
		fmt.Println("Migration error: ", err)
	} else {
		fmt.Println("Migrations applied successfully")
	}

	r, err := middleware.NewRBAC(conn)
	if err != nil {
		fmt.Println("Error while stablicing conection with postgres rbac. Err: ", err)
	}

	APIServer := api.NewApiServer(s.port, r.DB)
	APIServer.Run()
}

func (s *Server) Shutdown(ctx context.Context) error {
	if s.APIServer != nil {
		fmt.Println("Shotdown")
		return s.APIServer.Shutdown(ctx)
	}
	return nil
}
