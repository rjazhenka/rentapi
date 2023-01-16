package pgrent

import (
	"context"
	"github.com/rjazhenka/rentapi/pkg/api"
)

func (r *pgRentRepository) ModifySearch(ctx context.Context, req *api.ModifySearchRequest) (resp *api.ModifySearchResponse, er error) {
	sqlSrch := `
		insert into rent_search (params, chat_id, name)
		values ($1, $2, $3) 
		on conflict (name, chat_id) do update
			set name = $3, params = $1
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
