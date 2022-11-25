package repo

import (
	"context"
	"database/sql"
	"encoding/json"
)

type pgRentRepository struct {
	db *sql.DB
}

func (r *pgRentRepository) CreateRent(ctx context.Context, crRt *CreateRentDto) (rent *Rent, err error) {
	sqlSt := `
		insert into rent_turkey (title, rooms, price, country, city, region, district, description, link, source, images, imagesUrls)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		returning id
	`
	var id int64
	tgPhotos, _ := json.Marshal(crRt.TgPhotos)
	urlPhotos, _ := json.Marshal(crRt.UrlPhotos)
	err = r.db.QueryRow(sqlSt,
		crRt.Title,
		crRt.Rooms,
		crRt.Price,
		crRt.Country,
		crRt.City,
		crRt.Region,
		crRt.District,
		crRt.Description,
		crRt.Link,
		crRt.Source,
		tgPhotos,
		urlPhotos,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &Rent{
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
