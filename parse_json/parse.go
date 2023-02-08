package parsejson

import (
	"encoding/json"
	"log"
	"reflect"

	model "github.com/AlexandrYar/WBLO/models"
)

func Parsejson(data []byte) model.Order {
	var data_json map[string]interface{}
	var data_order = model.Order{}
	err := json.Unmarshal(data, &data_json)
	if err != nil {
		log.Println(err)
	}
	data_order.Order_uid = data_json["order_uid"].(string)
	data_order.Track_number = data_json["track_number"].(string)
	data_order.Entry = data_json["entry"].(string)
	data_order.Locale = data_json["locale"].(string)
	data_order.Internal_signature = data_json["internal_signature"].(string)
	data_order.Customer_id = data_json["customer_id"].(string)
	data_order.Delivery_service = data_json["delivery_service"].(string)
	data_order.Shardkey = data_json["shardkey"].(string)
	data_order.Sm_id = int(data_json["sm_id"].(float64))
	data_order.Date_created = data_json["date_created"].(string)
	data_order.Oof_shard = data_json["oof_shard"].(string)
	for i := range data_json {
		if reflect.TypeOf(data_json[i]) == reflect.TypeOf(data_json) {
			var data_item model.Item
			var data_payment model.Payment
			var data_delivery model.Delivery
			m := data_json[i].(map[string]interface{})
			switch i {
			case "items":
				data_item.Chrt_id = int(m["chrt_id"].(float64))
				data_item.Track_number = m["track_number"].(string)
				data_item.Price = int(m["price"].(float64))
				data_item.Rid = m["rid"].(string)
				data_item.Name = m["name"].(string)
				data_item.Sale = int(m["sale"].(float64))
				data_item.Size = m["size"].(string)
				data_item.Total_price = int(m["total_price"].(float64))
				data_item.Nm_id = int(m["nm_id"].(float64))
				data_item.Brand = m["brand"].(string)
				data_item.Status = int(m["status"].(float64))
				data_order.Items = data_item
			case "payment":
				data_payment.Transaction = m["transaction"].(string)
				data_payment.Request_id = m["request_id"].(string)
				data_payment.Currency = m["currency"].(string)
				data_payment.Provider = m["provider"].(string)
				data_payment.Amount = int(m["amount"].(float64))
				data_payment.Payment_dt = int64(m["payment_dt"].(float64))
				data_payment.Bank = m["bank"].(string)
				data_payment.Delivery_cost = int(m["delivery_cost"].(float64))
				data_payment.Goods_total = int(m["goods_total"].(float64))
				data_payment.Custom_fee = int(m["custom_fee"].(float64))
				data_order.Payment = data_payment
			case "delivery":
				data_delivery.Name = m["name"].(string)
				data_delivery.Phone = m["phone"].(string)
				data_delivery.Zip = m["zip"].(string)
				data_delivery.City = m["city"].(string)
				data_delivery.Address = m["address"].(string)
				data_delivery.Region = m["region"].(string)
				data_delivery.Email = m["email"].(string)
				data_order.Delivery = data_delivery
			}
		}

	}

	return data_order
}
