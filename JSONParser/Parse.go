package JSONParser

import (
	"KevinsProject/OrderStruct"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func ParseJSON() (OrderStruct.Orders, error) {
	jsonFile, err := os.Open("StoreOrders.json")
	var orders OrderStruct.Orders
	if err != nil {
		return orders, errors.New(fmt.Sprintf("There was an error opening the json. error = %s", err))
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Fatal("Was unable to close JSON File")
		}
	}(jsonFile)
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return orders, errors.New(fmt.Sprintf("There was an error closing the json. error = %s", err))
	}
	err = json.Unmarshal(byteValue, &orders)
	if err != nil {
		return orders, errors.New(fmt.Sprintf("There was an error decoding the json. error = %s", err))
	}
	return orders, nil
}
