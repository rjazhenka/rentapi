package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rjazhenka/rentapi/pkg/api"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"log"
	"strconv"
	"strings"
)

type pgRentRepository struct {
	db *sql.DB
}

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
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, 
		        $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, 
		        $23, $24, $25, $26, $27, $28, $29, $30, $31)
		on conflict(external_id) do update set
		                          title = EXCLUDED.title
		returning id
	`
	sqlOutbox := `
		insert into rent_turkey_outbox (id) values ($1) on conflict do nothing 
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
	}

	return &api.CreateRentResponse{
		Id: id,
	}, nil
}

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
			contact_tg_user_name
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

func NewPgRentRepository(db *sql.DB) *pgRentRepository {
	return &pgRentRepository{db: db}
}

func unwrapBool(b *wrapperspb.BoolValue) *bool {
	if b == nil {
		return nil
	}

	return &b.Value
}
