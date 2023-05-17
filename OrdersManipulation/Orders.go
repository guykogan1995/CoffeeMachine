package OrdersManipulation

import (
	"KevinsProject/JSONParser"
	"KevinsProject/OrderStruct"
	"errors"
	"fmt"
	"log"
	"sort"
	"strings"
)

var err error

type OrderArray OrderStruct.Orders

var Orders OrderStruct.Orders

func Parse() OrderStruct.Orders {
	Orders, err = JSONParser.ParseJSON()
	if err != nil {
		log.Fatal(err)
		return OrderStruct.Orders{}
	}
	return Orders
}

func GetOrders() OrderArray {
	Orders = Parse()
	return OrderArray(Orders)
}

func (ordersInput *OrderArray) GetOrderNames() (string, error) {
	str := ""
	if len(ordersInput.Orders) == 0 {
		return str, errors.New(fmt.Sprintf("There are no names to print orders empty. error = %s", err))
	}
	for _, order := range ordersInput.Orders {
		if order.Customer.FirstName == "" {
			break
		}
		str += fmt.Sprintf("ID: %d"+"\n\tCustomer name: %s\n\tTotal: $%s\n\tAddress: %s\n\tFullfillmentStatus: %s"+"\n", order.ID, order.Customer.FirstName+" "+order.Customer.LastName, order.TotalPrice, order.ShippingAddress.Address1, order.FulfillmentStatus)
	}
	return str, nil
}

func Filter(in *OrderArray, predicate func(order OrderStruct.Order) bool) OrderArray {
	var filtered OrderArray
	filtered = OrderArray{}
	for _, v := range in.Orders {
		if predicate(v) {
			filtered.Orders = append(filtered.Orders, v)
		}
	}
	return filtered
}

func ByUnFulfillment() func(order OrderStruct.Order) bool {
	return func(o OrderStruct.Order) bool {
		return o.FulfillmentStatus == "unfulfilled"
	}
}

func ByCustomerName(cName string) func(order OrderStruct.Order) bool {
	return func(o OrderStruct.Order) bool {
		return strings.Contains(strings.ToLower(o.Customer.FirstName+" "+o.Customer.LastName), strings.ToLower(cName))
	}
}

func ByItemName(item string) func(order OrderStruct.Order) bool {
	return func(o OrderStruct.Order) bool {
		return strings.Contains(strings.ToLower(o.LineItems[0].Name), strings.ToLower(item))
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
				return (*ordersInput).Orders[i].TotalPrice < (*ordersInput).Orders[j].TotalPrice
			})
		} else if upordown == "descending" {
			sort.Slice(*ordersInput, func(i, j int) bool {
				return (*ordersInput).Orders[i].TotalPrice > (*ordersInput).Orders[j].TotalPrice
			})
		}
		return *ordersInput
	case "date":
		if upordown == "ascending" {
			sort.Slice(*ordersInput, func(i, j int) bool {
				return (*ordersInput).Orders[i].CreatedAt.Before((*ordersInput).Orders[j].CreatedAt)
			})
		} else if upordown == "descending" {
			sort.Slice(*ordersInput, func(i, j int) bool {
				return (*ordersInput).Orders[i].CreatedAt.After((*ordersInput).Orders[j].CreatedAt)
			})
		}
		return *ordersInput
	case "customer name":
		if upordown == "ascending" {
			sort.Slice(*ordersInput, func(i, j int) bool {
				return (*ordersInput).Orders[i].Customer.FirstName <= (*ordersInput).Orders[j].Customer.FirstName
			})
		} else if upordown == "descending" {
			sort.Slice(*ordersInput, func(i, j int) bool {
				return (*ordersInput).Orders[i].Customer.FirstName >= (*ordersInput).Orders[j].Customer.FirstName
			})
		}
		return *ordersInput
	case "address":
		if upordown == "ascending" {
			sort.Slice(*ordersInput, func(i, j int) bool {
				return (*ordersInput).Orders[i].ShippingAddress.Address1 <= (*ordersInput).Orders[j].ShippingAddress.Address1
			})
		} else if upordown == "descending" {
			sort.Slice(*ordersInput, func(i, j int) bool {
				return (*ordersInput).Orders[i].ShippingAddress.Address1 >= (*ordersInput).Orders[j].ShippingAddress.Address1
			})
		}
		return *ordersInput
	default:
		return *ordersInput
	}
}

func (ordersInput *OrderArray) GetUnFulfilledOrders() OrderArray {
	return Filter(ordersInput, ByUnFulfillment())
}

func (ordersInput *OrderArray) GetFulfilledOrders() OrderArray {
	return Filter(ordersInput, ByFulfillment())
}

func (ordersInput *OrderArray) GetOrdersByName(cName string) OrderArray {
	return Filter(ordersInput, ByCustomerName(cName))
}

func (ordersInput *OrderArray) GetOrdersByItemName(item string) OrderArray {
	return Filter(ordersInput, ByItemName(item))
}

// ChangeStatus This function allows a user of the program
// To change the fulfillment status of an order to either
// "fulfilled" or "unfulfilled". Additionally, an order name
// is required to which order a user wants to change by the
// unique ticket name identifier
func (ordersInput *OrderArray) ChangeStatus(newStatus string, orderName int64) {
	isSet := false
	for i, order := range ordersInput.Orders {
		if order.ID == orderName {
			switch newStatus {
			case "fulfilled":
				(*ordersInput).Orders[i].FulfillmentStatus = "fulfilled"
				isSet = true
			case "unfulfilled":
				(*ordersInput).Orders[i].FulfillmentStatus = "unfulfilled"
				isSet = true
			default:
				log.Fatal("Command was not understood")
			}
			break
		}
	}
	if !isSet {
		log.Fatalf("No orderName of %d exists", orderName)
	}
}
