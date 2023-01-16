package pgrent

import (
	"database/sql"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type pgRentRepository struct {
	db *sql.DB
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
