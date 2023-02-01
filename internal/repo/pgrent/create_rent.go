package pgrent

import (
	"context"
	"encoding/json"
	"github.com/rjazhenka/rentapi/pkg/api"
)

func (r *pgRentRepository) CreateRent(ctx context.Context, crRt *api.CreateRentRequest) (rent *api.CreateRentResponse, err error) {
	sqlRent := `
		insert into rent_turkey (
		                         title, 
		                         rooms, rooms_label, 
		                         price, price_label, 
		                         country, country_label, 
		                         city, city_label, 
		                         region, region_label,
		                         district, district_label, 
		                         description, 
		                         link, 
		                         source, 
		                         images_tg_ids, 
		                         images_urls,
		                         address_label, address_elements,
		                         heating_gas_label, is_heating_gas,
		                         is_furnished,
		                         contact_label, contact,
		                         lat,
		                         long,
		                         external_id,
		                         tg_chat_id,
		                         tg_user_id,
		                         contact_tg_user_name
		                         )
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, INITCAP($11), $12, 
		        INITCAP($13), $14, $15, $16, $17, $18, $19, $20, $21, $22, 
		        $23, $24, $25, $26, $27, $28, $29, $30, $31)
		on conflict(external_id) do update set
		                          title = EXCLUDED.title
		returning id
	`
	sqlOutbox := `
		insert into rent_turkey_outbox (id) values ($1) on conflict do nothing 
	`

	sqlSrchQueue := `
			insert into rent_search_result_queue (search_id, rent_id, is_sent)
	select s.id
	    , ad.id
	    , false
	from rent_turkey ad
	join rent_search s on 
	   (s.params->>'max_price')::int >= ad.price
	   and (s.params->>'rooms')::int <= ad.rooms  
	   and ((s.params->>'towns_names')::jsonb ? ad.region_label or (s.params->>'towns_names')::jsonb = '[]'::jsonb or (s.params->>'towns_names') is null)
       and (((s.params->>'quarters_names')::jsonb ? ad.district_label) or (s.params->>'quarters_names')::jsonb = '[]'::jsonb or (s.params->>'quarter_names') is null)
	   and ad.id = $1;
	`

	var id int64
	tgPhotos, _ := json.Marshal(crRt.TgPhotos)
	urlPhotos, _ := json.Marshal(crRt.UrlPhotos)
	addressElements, _ := json.Marshal(crRt.AddressElements)

	err = r.db.QueryRow(sqlRent,
		crRt.Title,
		crRt.Rooms,
		crRt.RoomsLabel,
		crRt.Price,
		crRt.PriceLabel,
		crRt.Country,
		crRt.CountryLabel,
		crRt.City,
		crRt.CityLabel,
		crRt.Region,
		crRt.RegionLabel,
		crRt.District,
		crRt.DistrictLabel,
		crRt.Description,
		crRt.Link,
		crRt.Source,
		tgPhotos,
		urlPhotos,
		crRt.AddressLabel,
		addressElements,
		crRt.HeatingGasLabel,
		unwrapBool(crRt.GetIsHeatingGas()),
		unwrapBool(crRt.GetIsFurnished()),
		crRt.ContactLabel,
		crRt.Contact,
		crRt.GetLocation().GetLat(),
		crRt.GetLocation().GetLong(),
		crRt.ExternalId,
		crRt.TgChatId,
		crRt.TgUserId,
		crRt.ContactTgUserName,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	// will send only ads with filled rooms
	if crRt.RoomsLabel != "" {
		_, err = r.db.Exec(sqlOutbox, id)

		if err != nil {
			return nil, err
		}

		_, err = r.db.Exec(sqlSrchQueue, id)

		if err != nil {
			return nil, err
		}
	}

	return &api.CreateRentResponse{
		Id: id,
	}, nil
}
