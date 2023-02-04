package Model

type ModelDb struct {
	Host     string
	Port     string
	User     string
	Password string
	dbname   string
	sllmode  string
}

type Item struct {
	Chrt_id      int
	Track_number string
	Price        int
	Rid          string
	Name         string
	Sale         int
	Size         string
	Total_price  int
	Nm_id        int
	Brand        string
	Status       int
}

type Items []Item

type Payment struct {
	Transaction   string
	Request_id    string
	Currency      string // maybeee int ? 810 = USD
	Provider      string
	Amount        int
	Payment_dt    int64 // time t
	Bank          string
	Delivery_cost int
	Goods_total   int
	Custom_fee    int
}

type Delivery struct {
	Name    string
	Phone   string
	Zip     string
	City    string
	Address string
	Region  string
	Email   string
}

type Order struct {
	Order_uid          string
	Track_number       string
	Entry              string
	Delivery           Delivery
	Payment            Payment
	Items              Items
	Locale             string
	Internal_signature string
	Customer_id        string
	Delivery_service   string
	Shardkey           string
	Sm_id              int
	Date_created       string // time t nonono
	Oof_shard          string
}

type Orders []Order
