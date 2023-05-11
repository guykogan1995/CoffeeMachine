package main

import (
	"KevinsProject/OrdersManipulation"
	"fmt"
)

func main() {
	AllOrders := OrdersManipulation.GetOrders()
	UnFulfilledOrders := OrdersManipulation.GetUnFulfilledOrders()
	FulfilledOrders := OrdersManipulation.GetFulfilledOrders()
	customerOrders := OrdersManipulation.GetOrdersByName("Russell Winfield")
	AllOrders.SortBy("descending", "date")
	UnFulfilledOrders.SortBy("descending", "date")
	FulfilledOrders.SortBy("descending", "date")
	customerOrders.SortBy("descending", "date")

	fmt.Println("\nAll Orders: ")
	fmt.Println(AllOrders.GetOrderNames())
	fmt.Println("\nCustomer name orders: ")
	fmt.Println(customerOrders.GetOrderNames())
	fmt.Println("\n\nFulfilled Orders: ")
	fmt.Println(FulfilledOrders.GetOrderNames())
	fmt.Println("\nUnfulfilled Orders: ")
	fmt.Println(UnFulfilledOrders.GetOrderNames())

}
