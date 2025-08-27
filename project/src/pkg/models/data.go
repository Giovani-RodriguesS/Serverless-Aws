package models

import "encoding/json"

type Data struct {
	Type string          `json:"Type"`
	Data json.RawMessage `json:"Data"`
}

/*
{
  "type": "account",
  "data": {
    "id": "123",
    "name": "João Silva",
    "email": "joao@email.com"
  }
}
*/
