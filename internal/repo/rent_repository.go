package repo

import (
	"context"
)

type RentRepository interface {
	CreateRent(ctx context.Context, createRent *CreateRentDto) (rent *Rent, er error)
}

type CreateRentDto struct {
	Title       string
	Rooms       int32
	Price       float32
	Country     int32
	City        int32
	Region      int32
	District    int32
	Description string
	Link        string
	Source      int32
}

type Rent struct {
	Id          int64
	Title       string
	Rooms       int32
	Price       float32
	Country     int32
	City        int32
	Region      int32
	District    int32
	Description string
	Link        string
	Source      int32
}
