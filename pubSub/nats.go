package nats

import (
	"encoding/json"
	"log"

	model "github.com/AlexandrYar/WBLO/models"
	"github.com/nats-io/nats.go"
)

type PubSub struct {
}

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
	Set(Connection(), data_json)
}
