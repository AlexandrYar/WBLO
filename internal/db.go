package internal

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

func FindOrderById(conn *sql.DB, id int) model.Order {
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
