package models

import "encoding/json"

type Data struct {
	Type string          `json:"Type"`
	Data json.RawMessage `json:"Data"`
}

