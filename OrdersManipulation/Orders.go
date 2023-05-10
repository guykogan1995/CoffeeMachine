package OrdersManipulation

import (
	"KevinsProject/JSONParser"
	"KevinsProject/OrderStruct"
	"errors"
	"fmt"
	"log"
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
		str += order.Name + "\n"
	}
	return str, nil
}
