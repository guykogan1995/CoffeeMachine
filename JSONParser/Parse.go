package JSONParser

import (
	"KevinsProject/OrderStruct"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func ParseJSON() ([]OrderStruct.Order, error) {
	jsonFile, err := os.Open("Test_orders.json")
	var orders []OrderStruct.Order
	if err != nil {
		return orders, errors.New(fmt.Sprintf("There was an error opening the json. error = %s", err))
	}
	fmt.Println("Successfully Opened json file")
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return orders, errors.New(fmt.Sprintf("There was an error closing the json. error = %s", err))
	}
	err = json.Unmarshal(byteValue, &orders)
	if err != nil {
		return orders, errors.New(fmt.Sprintf("There was an error decoding the json. error = %s", err))
	}
	return orders, nil
}
