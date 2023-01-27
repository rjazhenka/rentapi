package pgrent

import (
	"context"
	"github.com/rjazhenka/rentapi/pkg/api"
)

func (r *pgRentRepository) ModifySearch(ctx context.Context, req *api.ModifySearchRequest) (resp *api.ModifySearchResponse, er error) {
	sqlSrch := `
		insert into rent_search (params, chat_id, name, state)
		values ($1, $2, $3, $4) 
		on conflict (name, chat_id) do update
			set name = $3, params = $1, state = $4
		returning id
	`
	params := &RentSearchParams{
		Rooms:         req.Rooms,
		MaxPrice:      req.MaxPrice,
		TownsIds:      req.TownsIds,
		TownsNames:    req.TownsNames,
		QuartersIds:   req.QuartersIds,
		QuartersNames: req.QuartersNames,
		isVnzh:        req.IsVnz,
	}

	var id int64
	err := r.db.QueryRow(sqlSrch, params, req.ChatId, req.Name, req.State).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &api.ModifySearchResponse{Id: id}, nil
}
