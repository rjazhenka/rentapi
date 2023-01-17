package repo

import (
	"context"
	"github.com/rjazhenka/rentapi/pkg/api"
)

type RentRepository interface {
	CreateRent(ctx context.Context, createRent *api.CreateRentRequest) (rent *api.CreateRentResponse, er error)
	GetRentToSend(ctx context.Context, req *api.GetRentToSendRequest) (resp *api.GetRentToSendResponse, er error)
	MarkAsSent(ctx context.Context, req *api.MarkAsSentRequest) (resp *api.MarkAsSentResponse, er error)
	CheckIfExist(ctx context.Context, req *api.CheckIfExistRequest) (resp *api.CheckIfExistResponse, er error)
	ModifySearch(ctx context.Context, req *api.ModifySearchRequest) (resp *api.ModifySearchResponse, er error)
	CreateSearch(ctx context.Context, req *api.CreateSearchRequest) (resp *api.CreateSearchResponse, er error)
	GetSearch(ctx context.Context, req *api.GetSearchRequest) (resp *api.GetSearchResponse, er error)
	GetSearchToSend(ctx context.Context, req *api.GetSearchToSendRequest) (resp *api.GetSearchToSendResponse, er error)
	MarkSearchAsSent(ctx context.Context, req *api.MarkSearchAsSentRequest) (resp *api.MarkSearchAsSentResponse, er error)
}
