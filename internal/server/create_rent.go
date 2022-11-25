package server

import (
	"context"
	"github.com/rjazhenka/rentapi/internal/repo"
	"github.com/rjazhenka/rentapi/pkg/api"
)

func (s *grpcServer) CreateRent(ctx context.Context, crReq *api.CreateRentRequest) (*api.CreateRentResponse, error) {
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
		TgPhotos:    crReq.TgPhotos,
		UrlPhotos:   crReq.UrlPhotos,
	}
	rent, err := s.rentRepo.CreateRent(ctx, crRentDto)
	if err != nil {
		return nil, err
	}
	return &api.CreateRentResponse{
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
		TgPhotos:    rent.TgPhotos,
		UrlPhotos:   rent.UrlPhotos,
	}, err
}
