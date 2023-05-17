package ShopifyAPI

import (
	"KevinsProject/OrderStruct"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func ShopifyRequest(endpoint string, accessToken string, method string, shopName string) []byte {
	url := fmt.Sprintf("https://%s.myshopify.com/admin/api/2023-04/%s.json?status=any", shopName, endpoint)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to begin the request to url: %s", url))
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Shopify-Access-Token", accessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable aquire response from: %s", req.URL))
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(fmt.Sprintf("Unable to close file"))
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to decipher response: %s", body))
	}
	return body
}

func GetDataForJSON(accessToken string) {
	var data OrderStruct.Orders
	body := ShopifyRequest(fmt.Sprintf("orders"), accessToken, "GET", "thefoundationcoffee")
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error: %s", err))
	}
	formattedResult, _ := json.MarshalIndent(data, "", "  ")
	for i, x := range data.Orders {
		fmt.Printf("Order %d: %v\n\n", i+1, x)
	}
	err = os.WriteFile("StoreOrders.json", formattedResult, 0644)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error: %s", err))
	}

}
