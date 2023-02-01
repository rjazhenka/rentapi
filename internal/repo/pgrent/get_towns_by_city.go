package pgrent

import (
	"context"
	"github.com/rjazhenka/rentapi/pkg/api"
	"log"
)

func (r *pgRentRepository) GetTownsByCity(_ context.Context, req *api.GetTownsByCityRequest) (resp *api.GetTownsByCityResponse, er error) {
	sql := `select 
    	id, 
    	name_tr 
from geo_town where city_id = $1
order by freq desc`

	rows, err := r.db.Query(sql, req.CityId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	resp = &api.GetTownsByCityResponse{}
	var town *api.Town

	for rows.Next() {
		town = &api.Town{}

		err := rows.Scan(
			&town.Id,
			&town.NameTr,
		)

		if err != nil {
			log.Println(err.Error())
		}

		resp.Towns = append(resp.Towns, town)
	}

	return resp, nil
}
