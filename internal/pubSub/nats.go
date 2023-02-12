package pubsub

import (
	"encoding/json"
	"log"

	"github.com/AlexandrYar/WBLO/internal/db"
	model "github.com/AlexandrYar/WBLO/models"
	parsejson "github.com/AlexandrYar/WBLO/parse_json"
	"github.com/nats-io/nats.go"
)

func Pub(order model.Order) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Panic(err.Error())
	}
	defer nc.Close()

	data, err := json.Marshal(&order)
	if err != nil {
		log.Panic(err.Error())
	}
	if err := nc.Publish("Order", data); err != nil {
		log.Panic(err.Error())
	}

	data_json := parsejson.Parsejson(data)
	db.Set(db.Connection(), data_json)
}
