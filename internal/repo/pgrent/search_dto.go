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
	Rooms         int32    `json:"rooms,omitempty"`
	MaxPrice      int32    `json:"max_price,omitempty"`
	TownsIds      []int32  `json:"towns_ids"`
	TownsNames    []string `json:"towns_names"`
	QuartersIds   []int32  `json:"quarters_ids"`
	QuartersNames []string `json:"quarters_names"`
	isVnzh        bool     `json:"is_vnz"`
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
