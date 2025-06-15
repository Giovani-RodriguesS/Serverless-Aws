package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type MyObject struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func ValidateJson(text string) error {

	trimmedText := strings.TrimSpace(text) // Remove espaço em branco

	if strings.HasPrefix(trimmedText, "{") && strings.HasSuffix(trimmedText, "}") {
		var jsonSimple MyObject

		decoderText := json.NewDecoder(bytes.NewReader([]byte(trimmedText)))
		decoderText.DisallowUnknownFields() // Não pode campos desconhecidos
		err := decoderText.Decode(&jsonSimple)

		if err != nil {
			fmt.Printf("Erro ao validar o json: %v", err)
		}
		
		// Faça a logica de mandar para a fila SQS 
		// Crie a logica de reportar erro ao Error Handler
		// Organize melhor a estrutura
	} else if strings.HasPrefix(trimmedText, "[") && strings.HasSuffix(trimmedText, "]") {
		var jsonList []MyObject
		
		decoderText := json.NewDecoder(bytes.NewReader([]byte(trimmedText)))
		decoderText.DisallowUnknownFields() // Não pode campos desconhecidos
		
		err := decoderText.Decode(&jsonList)

		if err != nil {
			fmt.Printf("Erro ao converter lista: %v", err)
		}
	
		fmt.Print(jsonList)
		
		return nil
	}


	return nil
}

func main (){
	//text := `{"name": "Item A", "value": 10}`
	text := `[{"name": "Item B", "value": 20}, {"name": "Item C", "value": 30}]`
	ValidateJson(text)
}
