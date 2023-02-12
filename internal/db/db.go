package db

import (
	"database/sql"
	"fmt"
	"log"

	model "github.com/AlexandrYar/WBLO/models"
	_ "github.com/lib/pq"
)

// Создание модели для подключения к БД
func NewDb() model.ModelDb {
	db := model.ModelDb{
		Host:     "localhost",
		Port:     "5432",
		User:     "wbuser",
		Password: "123",
		Dbname:   "wbdatabase",
		Sllmode:  "disable",
	}
	return db
}

// подключение к БД
func Connection() *sql.DB {
	db := NewDb()
	conString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", db.Host, db.Port, db.User, db.Password, db.Dbname, db.Sllmode)
	conn, err := sql.Open("postgres", conString)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func FindOrderById(conn *sql.DB, id string) model.Order {
	rows, err := conn.Query(`select "order_uid", "track_number", "entry", "locale", "internal_signature","customer_id","delivery_service","shardkey","sm_id","date_created","oof_shard" from orders where order_uid = $1`, id)
	if err != nil {
		log.Println(err)
	}
	var some_info model.Order
	for rows.Next() {

		err = rows.Scan(&some_info.Order_uid, &some_info.Track_number, &some_info.Entry, &some_info.Locale, &some_info.Internal_signature, &some_info.Customer_id, &some_info.Delivery_service, &some_info.Shardkey, &some_info.Sm_id, &some_info.Date_created, &some_info.Oof_shard)
		if err != nil {
			log.Println(err, "err 2")
		}
		log.Println("some info db:", some_info)
	}
	return some_info
}

func Set(conn *sql.DB, data model.Order) {
	err := conn.QueryRow(`INSERT INTO orders ("order_uid", "track_number", "entry", "locale", "internal_signature","customer_id","delivery_service","shardkey","sm_id","date_created","oof_shard") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, data.Order_uid, data.Track_number, data.Entry, data.Locale, data.Internal_signature, data.Customer_id, data.Delivery_service, data.Shardkey, data.Sm_id, data.Date_created, data.Oof_shard)
	if err != nil {
		log.Println(err)
	}
	err = conn.QueryRow(`INSERT INTO items ("chrt_id", "track_number", "price", "rid", "name", "sale", "size", "total_price", "nm_id", "brand", "status" ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, data.Items.Chrt_id, data.Track_number, data.Items.Price, data.Items.Rid, data.Items.Name, data.Items.Sale, data.Items.Size, data.Items.Total_price, data.Items.Nm_id, data.Items.Brand, data.Items.Status)
	if err != nil {
		log.Println(err)
	}
	err = conn.QueryRow(`INSERT INTO payment ( "transaction", "request_id", "currency", "provider", "amount", "paument_dt", "bank","delivery_cost", "goods_total", "custom_fee" ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`, data.Payment.Transaction, data.Payment.Request_id, data.Payment.Currency, data.Payment.Provider, data.Payment.Amount, data.Payment.Payment_dt, data.Payment.Bank, data.Payment.Delivery_cost, data.Payment.Goods_total, data.Payment.Custom_fee)
	if err != nil {
		log.Println(err)
	}
	log.Println("Пополение базы данных...")
}
