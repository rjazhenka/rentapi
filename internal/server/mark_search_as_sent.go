package server

import (
	"context"
	"github.com/rjazhenka/rentapi/pkg/api"
)

func (s *grpcServer) MarkSearchAsSent(ctx context.Context, req *api.MarkSearchAsSentRequest) (resp *api.MarkSearchAsSentResponse, er error) {
	resp, err := s.rentRepo.MarkSearchAsSent(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
