package server

import (
	"context"
	"github.com/rjazhenka/rentapi/pkg/api"
)

func (s *grpcServer) CreateSearch(ctx context.Context, req *api.CreateSearchRequest) (*api.CreateSearchResponse, error) {
	rent, err := s.rentRepo.CreateSearch(ctx, req)
	if err != nil {
		return nil, err
	}
	return rent, err
}
