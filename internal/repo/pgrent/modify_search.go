package pgrent

import (
	"context"
	"github.com/rjazhenka/rentapi/pkg/api"
	"golang.org/x/exp/slices"
)

func (r *pgRentRepository) ModifySearch(ctx context.Context, req *api.ModifySearchRequest) (resp *api.ModifySearchResponse, er error) {
	sqlSrch := `
		insert into rent_search (params, chat_id, name, state)
		values ($1, $2, $3, $4) 
		on conflict (name, chat_id) do update
			set name = $3, params = $1, state = $4
		returning id
	`
	towns := make([]SearchTown, len(req.Towns))
	townsNames := make([]string, len(req.Towns))
	var quartersNames []string

	hasLara := false //TODO refactor
	for i, t := range req.Towns {
		quarters := make([]SearchQuarter, len(t.Quarters))
		for k, q := range t.Quarters {
			if slices.Contains([]int{4851, 4853, 4854, 4855, 4856}, int(q.Id)) {
				hasLara = true
			}
			quartersNames = append(quartersNames, q.Name)
			quarters[k] = SearchQuarter{
				Id:   q.Id,
				Name: q.Name,
			}
		}
		towns[i] = SearchTown{
			t.Id,
			t.Name,
			quarters,
		}
		townsNames[i] = t.Name
	}
	if hasLara {
		quartersNames = append(quartersNames, "Lara")
	} else if i := slices.Index(quartersNames, "Lara"); i != -1 {
		quartersNames = slices.Delete(quartersNames, i, i+1)
	}
	params := &RentSearchParams{
		Rooms:         req.Rooms,
		MaxPrice:      req.MaxPrice,
		Towns:         towns,
		IsVnzh:        req.IsVnz,
		QuartersNames: quartersNames,
		TownsNames:    townsNames,
	}

	var id int64
	err := r.db.QueryRow(sqlSrch, params, req.ChatId, req.Name, req.State).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &api.ModifySearchResponse{Id: id}, nil
}
