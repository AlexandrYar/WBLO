package parsejson

import (
	"encoding/json"
	"log"

	model "github.com/AlexandrYar/WBLO/models"
)

func Parsejson(data []byte) model.Order {
	var data_json map[string]interface{}
	err := json.Unmarshal(data, data_json)
	if err != nil {
		log.Println(err)
	}
	log.Println(data_json)
}
