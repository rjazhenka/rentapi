package pgrent

import (
	"context"
	"github.com/rjazhenka/rentapi/pkg/api"
)

func (r *pgRentRepository) ModifySearch(ctx context.Context, req *api.ModifySearchRequest) (resp *api.ModifySearchResponse, er error) {
	sqlSrch := `
		update  rent_search set params = $1
		where chat_id = $2 and name = $3
		returning id
	`
	params := &RentSearchParams{
		Rooms:    req.Rooms,
		MaxPrice: req.MaxPrice,
	}

	var id int64
	err := r.db.QueryRow(sqlSrch, params, req.ChatId, req.Name).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &api.ModifySearchResponse{Id: id}, nil
}
