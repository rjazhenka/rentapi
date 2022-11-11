package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	v1 "rentapi/pkg/api"

	"rent_api/internal/repo"
	"rent_api/internal/server"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "realty"
)

func main() {
	db := getDbConn()
	defer db.Close()
	rentRepo := repo.NewPgRentRepository(db)

	s := grpc.NewServer(grpc.MaxSendMsgSize(10*10e6), grpc.MaxRecvMsgSize(10*10e6))
	reflection.Register(s)
	v1.RegisterRentServiceServer(s, server.NewGrpcServer(rentRepo))
	listener, err := net.Listen("tcp", "localhost:8080")
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
	return db
}
