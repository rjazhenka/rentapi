package repo

import (
	"context"
	"github.com/rjazhenka/rentapi/pkg/api"
)

type RentRepository interface {
	CreateRent(ctx context.Context, createRent *api.CreateRentRequest) (rent *api.CreateRentResponse, er error)
}
