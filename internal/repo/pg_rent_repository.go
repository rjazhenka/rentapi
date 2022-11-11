package repo

import (
	"context"
	"database/sql"
)

type pgRentRepository struct {
	db *sql.DB
}

func (r *pgRentRepository) CreateRent(ctx context.Context, crRt *CreateRentDto) (rent *Rent, err error) {
	sqlSt := `
		insert into rent_turkey (title, rooms, price, country, city, region, district, description, link, source)
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		returning id
	`
	var id int64
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
	}, nil
}

func NewPgRentRepository(db *sql.DB) *pgRentRepository {
	return &pgRentRepository{db: db}
}
