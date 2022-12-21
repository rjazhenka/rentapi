package server

import (
	"context"
	"github.com/rjazhenka/rentapi/pkg/api"
	"log"
)

func (s *grpcServer) GetRentToSend(ctx context.Context, req *api.GetRentToSendRequest) (resp *api.GetRentToSendResponse, er error) {
	resp, err := s.rentRepo.GetRentToSend(ctx, req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return resp, nil
}
