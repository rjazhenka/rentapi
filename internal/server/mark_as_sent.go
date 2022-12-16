package server

import (
	"context"
	"github.com/rjazhenka/rentapi/pkg/api"
)

func (s *grpcServer) MarkAsSent(ctx context.Context, req *api.MarkAsSentRequest) (resp *api.MarkAsSentResponse, er error) {
	resp, err := s.rentRepo.MarkAsSent(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
