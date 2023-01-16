package pgrent

import (
	"context"
	"errors"
	"fmt"
	"github.com/rjazhenka/rentapi/pkg/api"
	"strconv"
	"strings"
)

func (r *pgRentRepository) CheckIfExist(ctx context.Context, req *api.CheckIfExistRequest) (resp *api.CheckIfExistResponse, er error) {
	if len(req.Ids) == 0 {
		return nil, errors.New("ids can not be empty")
	}

	inSql := ""
	args := make([]any, 0, len(req.Ids))
	var lastI int
	for i, id := range req.Ids {
		lastI = i + 1
		inSql += "$" + strconv.Itoa(i+1) + ","
		args = append(args, any(id))
	}
	args = append(args, any(req.Source))
	inSql = strings.Trim(inSql, ",")

	sql := fmt.Sprintf("select external_id from rent_turkey where external_id in (%s) and source = $%d", inSql, lastI+1)

	stmt, err := r.db.Prepare(sql)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(args...)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	resp = &api.CheckIfExistResponse{}
	var id int64
	for rows.Next() {
		rows.Scan(&id)
		resp.Ids = append(resp.Ids, id)
	}

	return resp, nil
}
