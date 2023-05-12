package OrdersManipulation

import (
	"KevinsProject/JSONParser"
	"KevinsProject/OrderStruct"
	"errors"
	"fmt"
	"log"
	"sort"
)

var err error

type OrderArray []OrderStruct.Order

var Orders OrderArray

func Parse() OrderArray {
	Orders, err = JSONParser.ParseJSON()
	if err != nil {
		log.Fatal(err)
		return []OrderStruct.Order{}
	}
	return Orders
}

func GetOrders() OrderArray {
	Orders = Parse()
	return Orders
}

func (ordersInput *OrderArray) GetOrderNames() (string, error) {
	str := ""
	if len(Orders) == 0 {
		return str, errors.New(fmt.Sprintf("There are no names to print orders empty. error = %s", err))
	}
	for _, order := range *ordersInput {
		if order.Name == "" {
			break
		}
		str += fmt.Sprintf("ID: "+order.Name+"\n\tCustomer name: %s\n\tTotal: $%.2f\n\tAddress: %s\n\tFullfillmentStatus: %s"+"\n", order.ShippingName, order.Total, order.ShippingAddress1, order.FulfillmentStatus)
	}
	return str, nil
}

func Filter(in OrderArray, predicate func(order OrderStruct.Order) bool) OrderArray {
	var filtered OrderArray
	for _, v := range in {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func ByUnFulfillment() func(order OrderStruct.Order) bool {
	return func(o OrderStruct.Order) bool {
		return o.FulfillmentStatus == "unfulfilled" || o.FulfillmentStatus == ""
	}
}

func ByCustomerName(cName string) func(order OrderStruct.Order) bool {
	return func(o OrderStruct.Order) bool {
		return o.ShippingName == cName
	}
}

func ByFulfillment() func(order OrderStruct.Order) bool {
	return func(o OrderStruct.Order) bool {
		return o.FulfillmentStatus == "fulfilled"
	}
}

func (ordersInput *OrderArray) SortBy(upordown string, attribute string) OrderArray {
	switch attribute {
	case "total":
		if upordown == "ascending" {
			sort.Slice(*ordersInput, func(i, j int) bool {
				return (*ordersInput)[i].Total < (*ordersInput)[j].Total
			})
		} else if upordown == "descending" {
			sort.Slice(*ordersInput, func(i, j int) bool {
				return (*ordersInput)[i].Total > (*ordersInput)[j].Total
			})
		}
		return Orders
	case "date":
		if upordown == "ascending" {
			sort.Slice(*ordersInput, func(i, j int) bool {
				return (*ordersInput)[i].CreatedAt < (*ordersInput)[j].CreatedAt
			})
		} else if upordown == "descending" {
			sort.Slice(*ordersInput, func(i, j int) bool {
				return (*ordersInput)[i].CreatedAt > (*ordersInput)[j].CreatedAt
			})
		}
		return Orders
	case "customer name":
		if upordown == "ascending" {
			sort.Slice(*ordersInput, func(i, j int) bool {
				return (*ordersInput)[i].ShippingName <= (*ordersInput)[j].ShippingName
			})
		} else if upordown == "descending" {
			sort.Slice(*ordersInput, func(i, j int) bool {
				return (*ordersInput)[i].ShippingName >= (*ordersInput)[j].ShippingName
			})
		}
		return Orders
	case "address":
		if upordown == "ascending" {
			sort.Slice(*ordersInput, func(i, j int) bool {
				return (*ordersInput)[i].ShippingAddress1 <= (*ordersInput)[j].ShippingAddress1
			})
		} else if upordown == "descending" {
			sort.Slice(*ordersInput, func(i, j int) bool {
				return (*ordersInput)[i].ShippingAddress1 >= (*ordersInput)[j].ShippingAddress1
			})
		}
		return Orders
	default:
		return Orders
	}
}

func GetUnFulfilledOrders() OrderArray {
	return Filter(Orders, ByUnFulfillment())
}

func GetFulfilledOrders() OrderArray {
	return Filter(Orders, ByFulfillment())
}

func GetOrdersByName(cName string) OrderArray {
	return Filter(Orders, ByCustomerName(cName))
}

// ChangeStatus This function allows a user of the program
// To change the fulfillment status of an order to either
// "fulfilled" or "unfulfilled". Additionally, an order name
// is required to which order a user wants to change by the
// unique ticket name identifier
func (ordersInput *OrderArray) ChangeStatus(newStatus string, orderName string) {
	isSet := false
	for i, order := range *ordersInput {
		if order.Name == orderName {
			switch newStatus {
			case "fulfilled":
				(*ordersInput)[i].FulfillmentStatus = "fulfilled"
				isSet = true
			case "unfulfilled":
				(*ordersInput)[i].FulfillmentStatus = "unfulfilled"
				isSet = true
			default:
				log.Fatal("Command was not understood")
			}
			break
		}
	}
	if !isSet {
		log.Fatalf("No orderName of %s exists", orderName)
	}
}
