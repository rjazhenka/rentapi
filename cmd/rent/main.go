package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rjazhenka/rentapi/internal/repo"
	"github.com/rjazhenka/rentapi/internal/server"
	"github.com/rjazhenka/rentapi/pkg/api"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	err := initEnv()
	if err != nil {
		panic("Error config initialization")
	}

	db := getDbConn()
	defer db.Close()
	rentRepo := repo.NewPgRentRepository(db)

	s := grpc.NewServer()
	api.RegisterRentServiceServer(s, server.NewGrpcServer(rentRepo))
	listener, err := net.Listen("tcp", ":"+viper.GetString("SERVER_PORT"))
	log.Printf("Listening on port :%s", viper.GetString("SERVER_PORT"))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initEnv() error {
	env := os.Getenv("RENT_ENV")
	if env == "" {
		env = "local"
	}
	log.Printf("Start app with env %s", env)
	viper.SetConfigFile(".env/" + env + ".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return err
}

func getDbConn() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("POSTGRES_HOST"),
		viper.GetString("POSTGRES_PORT"),
		viper.GetString("POSTGRES_USER"),
		viper.GetString("POSTGRES_PASSWORD"),
		viper.GetString("POSTGRES_DBNAME"),
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	log.Println("Connection established")
	return db
}
