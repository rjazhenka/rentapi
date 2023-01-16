package pgrent

import (
	"context"
	"errors"
	"github.com/rjazhenka/rentapi/pkg/api"
)

func (r *pgRentRepository) MarkAsSent(ctx context.Context, req *api.MarkAsSentRequest) (resp *api.MarkAsSentResponse, er error) {
	if len(req.Items) == 0 {
		return nil, errors.New("ids can not be empty")
	}
	sql := "update rent_turkey_outbox set is_sent = true, sent_at = NOW(), tg_message_id = $1, tg_message_desc_id = $2 where id = $3"
	stmt, err := r.db.Prepare(sql)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for _, item := range req.Items {

		rows, err := stmt.Query(item.TgMessageId, item.TgMessageDescId, item.Id)

		if err != nil {
			return nil, err
		}
		rows.Close()
	}
	return &api.MarkAsSentResponse{}, nil
}
