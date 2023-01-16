package server

import (
	"context"
	"github.com/rjazhenka/rentapi/pkg/api"
)

func (s *grpcServer) ModifySearch(ctx context.Context, req *api.ModifySearchRequest) (*api.ModifySearchResponse, error) {
	rent, err := s.rentRepo.ModifySearch(ctx, req)
	if err != nil {
		return nil, err
	}
	return rent, err
}
