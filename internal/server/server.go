package server

import (
	"github.com/rjazhenka/rentapi/internal/repo"
	"github.com/rjazhenka/rentapi/pkg/api"
)

type grpcServer struct {
	api.UnimplementedRentServiceServer
	rentRepo repo.RentRepository
}

func NewGrpcServer(rentRepo repo.RentRepository) *grpcServer {
	s := &grpcServer{
		rentRepo: rentRepo,
	}
	return s
}
