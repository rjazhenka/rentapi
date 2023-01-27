package pgrent

import (
	"context"
	"github.com/rjazhenka/rentapi/pkg/api"
)

func (r *pgRentRepository) GetSearch(ctx context.Context, req *api.GetSearchRequest) (resp *api.GetSearchResponse, er error) {
	sqlSrch := "select id, name, chat_id, params, state from rent_search where chat_id = $1 and name = $2 limit 1"
	row := r.db.QueryRow(sqlSrch, req.ChatId, req.Name)

	srch := &RentSearch{}
	err := row.Scan(&srch.Id, &srch.Name, &srch.ChatId, &srch.Params, &srch.State)

	if err != nil {
		return nil, err
	}

	return &api.GetSearchResponse{
		Id:            srch.Id,
		Rooms:         srch.Params.Rooms,
		MaxPrice:      srch.Params.MaxPrice,
		TownsNames:    srch.Params.TownsNames,
		TownsIds:      srch.Params.TownsIds,
		QuartersNames: srch.Params.QuartersNames,
		QuartersIds:   srch.Params.QuartersIds,
		IsVnz:         srch.Params.isVnzh,
		ChatId:        srch.ChatId,
		Name:          srch.Name,
		State:         srch.State,
	}, nil
}
