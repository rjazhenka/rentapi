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

	towns := make([]*api.SearchTown, len(srch.Params.Towns))

	for i, t := range srch.Params.Towns {
		quarters := make([]*api.SearchQuarter, len(t.Quarters))
		for k, q := range t.Quarters {
			quarters[k] = &api.SearchQuarter{
				Id:   q.Id,
				Name: q.Name,
			}
		}
		town := &api.SearchTown{
			Id:       t.Id,
			Name:     t.Name,
			Quarters: quarters,
		}
		towns[i] = town
	}

	return &api.GetSearchResponse{
		Id:       srch.Id,
		Rooms:    srch.Params.Rooms,
		MaxPrice: srch.Params.MaxPrice,
		IsVnz:    srch.Params.IsVnzh,
		ChatId:   srch.ChatId,
		Name:     srch.Name,
		State:    srch.State,
		Towns:    towns,
	}, nil
}
