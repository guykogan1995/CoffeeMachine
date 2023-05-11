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
		str += fmt.Sprintf(order.Name+" Total: $%.2f"+"\n", order.Total)
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
