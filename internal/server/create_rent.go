package server

import (
	"context"
	v1 "rentapi/pkg/api"

	"rent_api/internal/repo"
)

func (s *grpcServer) CreateRent(ctx context.Context, crReq *v1.CreateRentRequest) (*v1.CreateRentResponse, error) {
	crRentDto := &repo.CreateRentDto{
		Title:       crReq.Title,
		Rooms:       crReq.Rooms,
		Price:       crReq.Price,
		Country:     crReq.Country,
		City:        crReq.City,
		Region:      crReq.Region,
		District:    crReq.District,
		Description: crReq.Description,
		Link:        crReq.Link,
		Source:      crReq.Source,
	}
	rent, err := s.rentRepo.CreateRent(ctx, crRentDto)
	if err != nil {
		return nil, err
	}
	return &v1.CreateRentResponse{
		Id:          rent.Id,
		Title:       rent.Title,
		Rooms:       rent.Rooms,
		Price:       rent.Price,
		Country:     rent.Country,
		City:        rent.City,
		Region:      rent.Region,
		District:    rent.District,
		Description: rent.Description,
		Link:        rent.Link,
		Source:      rent.Source,
	}, err
}
