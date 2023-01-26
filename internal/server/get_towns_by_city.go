package server

import (
	"context"
	"github.com/rjazhenka/rentapi/pkg/api"
	"log"
)

func (s *grpcServer) GetTownsByCity(ctx context.Context, req *api.GetTownsByCityRequest) (resp *api.GetTownsByCityResponse, er error) {
	resp, err := s.rentRepo.GetTownsByCity(ctx, req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return resp, nil
}
