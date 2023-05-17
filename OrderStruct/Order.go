package OrderStruct

import "time"

type Orders struct {
	Orders []Order `json:"orders"`
}

type Order struct {
	AdminGraphqlApiId string     `json:"admin_graphql_api_id"`
	AppId             int64      `json:"app_id"`
	BillingAddress    Address    `json:"billing_address"`
	ShippingAddress   Address    `json:"shipping_address"`
	Customer          Customer   `json:"customer"`
	LineItems         []LineItem `json:"line_items"`
	FulfillmentStatus string     `json:"fulfillment_status"`
	ID                int64      `json:"id"`
	TotalPrice        string     `json:"total_price"`
	CreatedAt         time.Time  `json:"created_at"`
	// Add more fields as needed
}

type Address struct {
	Address1  string `json:"address1"`
	City      string `json:"city"`
	Country   string `json:"country"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Name      string `json:"name"`
	Zip       string `json:"zip"`
	// Add more fields as needed
}

type Customer struct {
	AcceptsMarketing bool   `json:"accepts_marketing"`
	Email            string `json:"email"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	// Add more fields as needed
}

type LineItem struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	SKU      string `json:"sku"`
	Quantity int    `json:"quantity"`
	// Add more fields as needed
}
