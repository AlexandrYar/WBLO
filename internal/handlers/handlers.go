package handlers

import (
	"io/ioutil"
	"log"

	"github.com/AlexandrYar/WBLO/internal/db"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	id := c.PostForm("id")
	order_info := db.FindOrderById(db.Connection(), id)
	log.Println(order_info, " ", id)
	c.HTML(200, "index.html", gin.H{
		"Order_uid":          order_info.Order_uid,
		"Track_number":       order_info.Track_number,
		"Entry":              order_info.Entry,
		"Locale":             order_info.Locale,
		"Internal_signature": order_info.Internal_signature,
		"Customer_id":        order_info.Customer_id,
		"Delivery_service":   order_info.Delivery_service,
		"Shardkey":           order_info.Shardkey,
		"Sm_id":              order_info.Sm_id,
		"Date_created":       order_info.Date_created,
		"Oof_shard":          order_info.Oof_shard,
	})
}

func Json(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
	}

}
