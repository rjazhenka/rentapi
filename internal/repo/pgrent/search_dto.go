package pgrent

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type RentSearch struct {
	Id     int64
	Name   string
	ChatId int64
	Params *RentSearchParams
	State  int32
}

type RentSearchParams struct {
	Rooms         int32        `json:"rooms,omitempty"`
	MaxPrice      int32        `json:"max_price,omitempty"`
	Towns         []SearchTown `json:"towns"`
	IsVnzh        bool         `json:"is_vnz"`
	TownsNames    []string     `json:"towns_names"`
	QuartersNames []string     `json:"quarters_names"`
}

type SearchTown struct {
	Id       int32           `json:"id"`
	Name     string          `json:"name"`
	Quarters []SearchQuarter `json:"quarters"`
}

type SearchQuarter struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

func (s RentSearchParams) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *RentSearchParams) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &s)
}
