package main

import (
	"KevinsProject/OrdersManipulation"
	"fmt"
)

func main() {
	Orders := OrdersManipulation.GetOrders()
	fmt.Println(Orders.GetOrderNames())
}
