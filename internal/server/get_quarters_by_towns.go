package server

import (
	"context"
	"github.com/rjazhenka/rentapi/pkg/api"
	"log"
)

func (s *grpcServer) GetQuartersByTowns(ctx context.Context, req *api.GetQuartersByTownsRequest) (resp *api.GetQuartersByTownsResponse, er error) {
	resp, err := s.rentRepo.GetQuartersByTowns(ctx, req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return resp, nil
}
