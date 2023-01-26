package pgrent

import (
	"context"
	"github.com/lib/pq"
	"github.com/rjazhenka/rentapi/pkg/api"
	"log"
)

func (r *pgRentRepository) GetDistrictsByTowns(_ context.Context, req *api.GetQuartersByTownsRequest) (resp *api.GetQuartersByTownsResponse, er error) {
	sql := `select 
    	q.id quarter_id, 
    	q.name_tr quarter_name_tr,
    	is_vnz,
    	q.town_id town_id,
    	t.name_tr town_name
from geo_quarter q
join geo_town t on q.town_id = t.id
where q.town_id = any($1)`

	rows, err := r.db.Query(sql, pq.Array(req.TownIds))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	resp = &api.GetQuartersByTownsResponse{}
	var quarter *api.Quarter

	for rows.Next() {
		quarter = &api.Quarter{
			Town: &api.Town{},
		}

		err := rows.Scan(
			&quarter.Id,
			&quarter.NameTr,
			&quarter.IsVnzh,
			&quarter.Town.Id,
			&quarter.Town.NameTr,
		)

		if err != nil {
			log.Println(err.Error())
		}

		resp.Quarters = append(resp.Quarters, quarter)
	}

	return resp, nil
}
