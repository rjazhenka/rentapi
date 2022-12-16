package server

import (
	"context"
	"github.com/rjazhenka/rentapi/pkg/api"
)

func (s *grpcServer) CheckIfExist(ctx context.Context, req *api.CheckIfExistRequest) (*api.CheckIfExistResponse, error) {
	res, err := s.rentRepo.CheckIfExist(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, err
}
