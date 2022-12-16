package repo

import (
	"context"
	"github.com/rjazhenka/rentapi/pkg/api"
)

type RentRepository interface {
	CreateRent(ctx context.Context, createRent *api.CreateRentRequest) (rent *api.CreateRentResponse, er error)
	GetRentToSend(ctx context.Context, req *api.GetRentToSendRequest) (resp *api.GetRentToSendResponse, er error)
	MarkAsSent(ctx context.Context, req *api.MarkAsSentRequest) (resp *api.MarkAsSentResponse, er error)
}
