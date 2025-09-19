package internal

import (
	"encoding/json"
	"fmt"
	"github.com/Giovani-RodriguesS/Serverless-Aws/project/src/pkg/models"
)

func ParseJsonToItem(body string) (models.Data, error) {
	// Converte o dado bruto para ser possivel identificar o objeto
	var data models.Data
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		return models.Data{}, err
	}
	return data, nil
}

func WrapUpItem(data models.Data) (any, error) {
	// Identifica qual item será gravado no Dynamo
	switch data.Type {
	case "account":
		var acc models.Account
		if err := json.Unmarshal(data.Data, &acc); err != nil {
			return nil, err
		}

		return &acc, nil

	case "transaction":
		var tsc models.Transaction
		if err := json.Unmarshal(data.Data, &tsc); err != nil {
			return nil, err
		}
		return &tsc, nil

	default:
		return nil, fmt.Errorf("tipo não reconhecido: %s", data.Type)
	}
}
