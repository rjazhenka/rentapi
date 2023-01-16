package server

import (
	"context"
	"github.com/rjazhenka/rentapi/pkg/api"
)

func (s *grpcServer) GetSearch(ctx context.Context, req *api.GetSearchRequest) (*api.GetSearchResponse, error) {
	res, err := s.rentRepo.GetSearch(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, err
}
