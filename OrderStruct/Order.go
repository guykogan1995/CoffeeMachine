package OrderStruct

type Order struct {
	Name                      string  `json:"Name"`
	Email                     string  `json:"Email"`
	FinancialStatus           string  `json:"Financial Status"`
	PaidAt                    string  `json:"Paid at"`
	FulfillmentStatus         string  `json:"Fulfillment Status"`
	FulfilledAt               string  `json:"Fulfilled at"`
	AcceptsMarketing          string  `json:"Accepts Marketing"`
	Currency                  string  `json:"Currency"`
	Subtotal                  float64 `json:"Subtotal"`
	Shipping                  float64 `json:"Shipping"`
	Taxes                     float64 `json:"Taxes"`
	Total                     float64 `json:"Total"`
	DiscountCode              string  `json:"Discount Code"`
	DiscountAmount            float64 `json:"Discount Amount"`
	ShippingMethod            string  `json:"Shipping Method"`
	CreatedAt                 string  `json:"Created at"`
	LineitemQuantity          float64 `json:"Lineitem quantity"`
	LineitemName              string  `json:"Lineitem name"`
	LineitemPrice             float64 `json:"Lineitem price"`
	LineitemCompareAtPrice    string  `json:"Lineitem compare at price"`
	LineitemSku               string  `json:"Lineitem sku"`
	LineitemRequiresShipping  bool    `json:"Lineitem requires shipping"`
	LineitemTaxable           bool    `json:"Lineitem taxable"`
	LineitemFulfillmentStatus string  `json:"Lineitem fulfillment status"`
	BillingName               string  `json:"Billing Name"`
	BillingStreet             string  `json:"Billing Street"`
	BillingAddress1           string  `json:"Billing Address1"`
	BillingAddress2           string  `json:"Billing Address2"`
	BillingCompany            string  `json:"Billing Company"`
	BillingCity               string  `json:"Billing City"`
	BillingZip                int     `json:"Billing Zip"`
	BillingProvince           string  `json:"Billing Province"`
	BillingCountry            string  `json:"Billing Country"`
	BillingPhone              int     `json:"Billing Phone"`
	ShippingName              string  `json:"Shipping Name"`
	ShippingStreet            string  `json:"Shipping Street"`
	ShippingAddress1          string  `json:"Shipping Address1"`
	ShippingAddress2          string  `json:"Shipping Address2"`
	ShippingCompany           string  `json:"Shipping Company"`
	ShippingCity              string  `json:"Shipping City"`
	ShippingZip               int     `json:"Shipping Zip"`
	ShippingProvince          string  `json:"Shipping Province"`
	ShippingCountry           string  `json:"Shipping Country"`
	ShippingPhone             int     `json:"Shipping Phone"`
	Notes                     string  `json:"Notes"`
	NoteAttributes            string  `json:"Note Attributes"`
	CancelledAt               string  `json:"Cancelled at"`
	PaymentMethod             string  `json:"Payment Method"`
	PaymentReference          string  `json:"Payment Reference"`
	RefundedAmount            float64 `json:"Refunded Amount"`
	Vendor                    string  `json:"Vendor"`
	OutstandingBalance        float64 `json:"Outstanding Balance"`
	Employee                  string  `json:"Employee"`
	Location                  string  `json:"Location"`
	DeviceID                  string  `json:"Device ID"`
	Id                        float64 `json:"Id"`
	Tags                      string  `json:"Tags"`
	RiskLevel                 string  `json:"Risk Level"`
	Source                    float64 `json:"Source"`
	LineitemDiscount          float64 `json:"Lineitem discount"`
	Tax1Name                  string  `json:"Tax 1 Name"`
	Tax1Value                 string  `json:"Tax 1 Value"`
	Tax2Name                  string  `json:"Tax 2 Name"`
	Tax2Value                 string  `json:"Tax 2 Value"`
	Tax3Name                  string  `json:"Tax 3 Name"`
	Tax3Value                 string  `json:"Tax 3 Value"`
	Tax4Name                  string  `json:"Tax 4 Name"`
	Tax4Value                 string  `json:"Tax 4 Value"`
	Tax5Name                  string  `json:"Tax 5 Name"`
	Tax5Value                 string  `json:"Tax 5 Value"`
	Phone                     string  `json:"Phone"`
	ReceiptNumber             string  `json:"Receipt Number"`
	Duties                    string  `json:"Duties"`
	BillingProvinceName       string  `json:"Billing Province Name"`
	ShippingProvinceName      string  `json:"Shipping Province Name"`
	PaymentID                 string  `json:"Payment ID"`
	PaymentTermsName          string  `json:"Payment Terms Name"`
	NextPaymentDueAt          string  `json:"Next Payment Due At"`
	PaymentReferences         string  `json:"Payment References"`
}
