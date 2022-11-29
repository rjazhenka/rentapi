package server

import (
	"context"
	"github.com/rjazhenka/rentapi/pkg/api"
)

func (s *grpcServer) CreateRent(ctx context.Context, crReq *api.CreateRentRequest) (*api.CreateRentResponse, error) {
	rent, err := s.rentRepo.CreateRent(ctx, crReq)
	if err != nil {
		return nil, err
	}
	return rent, err
}
