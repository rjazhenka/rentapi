package pgrent

import (
	"context"
	"github.com/rjazhenka/rentapi/pkg/api"
)

func (r *pgRentRepository) CreateSearch(ctx context.Context, req *api.CreateSearchRequest) (resp *api.CreateSearchResponse, er error) {
	sqlSrch := `
		insert into rent_search (chat_id, params, name)
		values ($1, $2, $3)
		returning id
	`
	params := &RentSearchParams{
		Rooms:    req.Rooms,
		MaxPrice: req.MaxPrice,
	}

	var id int64
	err := r.db.QueryRow(sqlSrch, req.ChatId, params, req.Name).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &api.CreateSearchResponse{Id: id}, nil
}
