package pgrent

import (
	"context"
	"errors"
	"github.com/rjazhenka/rentapi/pkg/api"
)

func (r *pgRentRepository) MarkSearchAsSent(ctx context.Context, req *api.MarkSearchAsSentRequest) (resp *api.MarkSearchAsSentResponse, er error) {
	if len(req.Items) == 0 {
		return nil, errors.New("ids can not be empty")
	}
	sql := "update rent_search_result_queue set is_sent = true, sent_at = NOW() where id = $1"
	stmt, err := r.db.Prepare(sql)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for _, item := range req.Items {
		rows, err := stmt.Query(item.Id)

		if err != nil {
			return nil, err
		}
		rows.Close()
	}
	return &api.MarkSearchAsSentResponse{}, nil
}
