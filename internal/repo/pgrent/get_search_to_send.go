package pgrent

import (
	"context"
	"encoding/json"
	"github.com/rjazhenka/rentapi/pkg/api"
	"log"
)

func (r *pgRentRepository) GetSearchToSend(ctx context.Context, req *api.GetSearchToSendRequest) (resp *api.GetSearchToSendResponse, er error) {
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
			o.tg_message_id, -- TODO refactor
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
			coalesce(r.contact_tg_user_name, ''),
			s.chat_id,
			q.id
		from rent_turkey r
		join public.rent_turkey_outbox o on o.id = r.id and o.is_sent = true -- TODO refactor. 
		join rent_search_result_queue q on r.id = q.rent_id and q.is_sent = false
		join rent_search s on q.search_id = s.id
		order by r.id
		for update skip locked
		limit $1`
	rows, err := r.db.Query(sql, req.Limit)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	resp = &api.GetSearchToSendResponse{}

	for rows.Next() {
		ad := &api.GetRentToSendResponseItem{
			Location: &api.Location{},
		}
		var chatId, id int64

		var tgImages, urlImages string
		err := rows.Scan(&ad.Id,
			&ad.Title,
			&ad.RoomsLabel,
			&ad.Price,
			&ad.CityLabel,
			&ad.RegionLabel,
			&ad.DistrictLabel,
			&ad.Description,
			&ad.Link,
			&ad.Source,
			&tgImages,
			&urlImages,
			&ad.AddressLabel,
			&ad.ContactLabel,
			&ad.Contact,
			&ad.ExternalId,
			&ad.TgChatId,
			&ad.TgUserId,
			&ad.Location.Lat,
			&ad.Location.Long,
			&ad.HeatingGasLabel,
			&ad.HasHeatiing,
			&ad.ContactTgUserName,
			&chatId,
			&id,
		)

		if err != nil {
			log.Println(err.Error())
		}

		item := &api.GetSearchToSendResponseItem{
			Ad:     ad,
			ChatId: chatId,
			Id:     id,
		}

		json.Unmarshal([]byte(tgImages), &ad.TgPhotos)
		json.Unmarshal([]byte(urlImages), &ad.UrlPhotos)
		resp.Items = append(resp.Items, item)
	}

	return resp, nil
}
