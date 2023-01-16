package pgrent

import (
	"context"
	"encoding/json"
	"github.com/rjazhenka/rentapi/pkg/api"
	"log"
)

func (r *pgRentRepository) GetRentToSend(ctx context.Context, req *api.GetRentToSendRequest) (resp *api.GetRentToSendResponse, er error) {
	sql := `select 
			r.id,
			title,
			rooms_label,
			price,
			city_label,
			region_label,
			district_label,
			description,
			link,
			source,
			images_tg_ids,
			images_urls,
			address_label,
			contact_label,
			contact,
			external_id,
			tg_chat_id,
			tg_user_id,
			lat,
			long,
			heating_gas_label,
			coalesce(is_heating_gas, false),
			r.contact_tg_user_name
		from rent_turkey r
		join rent_turkey_outbox o on r.id = o.id and o.is_sent = false
		order by r.id
		for update skip locked
		limit $1`
	rows, err := r.db.Query(sql, req.Limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	resp = &api.GetRentToSendResponse{}

	var item *api.GetRentToSendResponseItem
	for rows.Next() {
		item = &api.GetRentToSendResponseItem{Location: &api.Location{}}
		var tgImages, urlImages string
		err := rows.Scan(&item.Id,
			&item.Title,
			&item.RoomsLabel,
			&item.Price,
			&item.CityLabel,
			&item.RegionLabel,
			&item.DistrictLabel,
			&item.Description,
			&item.Link,
			&item.Source,
			&tgImages,
			&urlImages,
			&item.AddressLabel,
			&item.ContactLabel,
			&item.Contact,
			&item.ExternalId,
			&item.TgChatId,
			&item.TgUserId,
			&item.Location.Lat,
			&item.Location.Long,
			&item.HeatingGasLabel,
			&item.HasHeatiing,
			&item.ContactTgUserName,
		)

		if err != nil {
			log.Println(err.Error())
		}

		json.Unmarshal([]byte(tgImages), &item.TgPhotos)
		json.Unmarshal([]byte(urlImages), &item.UrlPhotos)
		resp.Items = append(resp.Items, item)
	}

	return resp, nil
}
