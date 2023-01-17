package server

import (
	"context"
	"github.com/rjazhenka/rentapi/pkg/api"
	"log"
)

func (s *grpcServer) GetSearchToSend(ctx context.Context, req *api.GetSearchToSendRequest) (resp *api.GetSearchToSendResponse, er error) {
	resp, err := s.rentRepo.GetSearchToSend(ctx, req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return resp, nil
}
