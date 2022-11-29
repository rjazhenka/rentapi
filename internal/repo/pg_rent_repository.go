package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/rjazhenka/rentapi/pkg/api"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type pgRentRepository struct {
	db *sql.DB
}

func (r *pgRentRepository) CreateRent(ctx context.Context, crRt *api.CreateRentRequest) (rent *api.CreateRentResponse, err error) {
	sqlSt := `
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
		                         external_id
		                         )
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, 
		        $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, 
		        $23, $24, $25, $26, $27, $28)
		returning id
	`
	var id int64
	tgPhotos, _ := json.Marshal(crRt.TgPhotos)
	urlPhotos, _ := json.Marshal(crRt.UrlPhotos)
	addressElements, _ := json.Marshal(crRt.AddressElements)

	err = r.db.QueryRow(sqlSt,
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
		crRt.GetLocation().GetLat(),
		crRt.ExternalId,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &api.CreateRentResponse{
		Id:          id,
		Title:       crRt.Title,
		Rooms:       crRt.Rooms,
		Price:       crRt.Price,
		Country:     crRt.Country,
		City:        crRt.City,
		Region:      crRt.Region,
		District:    crRt.District,
		Description: crRt.Description,
		Link:        crRt.Link,
		Source:      crRt.Source,
		TgPhotos:    crRt.TgPhotos,
	}, nil
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
