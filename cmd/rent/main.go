package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rjazhenka/rentapi/internal/repo"
	"github.com/rjazhenka/rentapi/internal/server"
	"github.com/rjazhenka/rentapi/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "realty"
)

func main() {
	db := getDbConn()
	defer db.Close()
	rentRepo := repo.NewPgRentRepository(db)

	s := grpc.NewServer()
	api.RegisterRentServiceServer(s, server.NewGrpcServer(rentRepo))
	listener, err := net.Listen("tcp", ":8080")
	log.Printf("Listen to the port %d", 8080)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getDbConn() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	log.Println("Connection established")
	return db
}
