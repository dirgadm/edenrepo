package edenrepo

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// CoreGateway ...
type CoreGateway struct {
	Client http.Client
}

// Call : Base Method to call Core API
func (client *CoreGateway) Call(method, path string, body io.Reader, v interface{}) error {
	credentials := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpveUBlZGVuZmFybS5pZCIsInBhc3N3b3JkIjoiZGV2MTIzIn0.tgiwi4Mwcj5l2oyrCKIpN6BE7dHCO18vqBAnfk1Q4hU"
	var bearer = "Bearer " + credentials

	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Authorization", bearer)
	resp, err := client.Client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	defer resp.Body.Close()

	return err
}

// GetStockTransferByID ...
func (client *CoreGateway) GetStockTransferByID(id string) (http.Response, error) {
	resp := http.Response{}

	url := baseURL + "/v1/inventory/stock-transfer/" + id + ""
	err := client.Call("GET", url, nil, &resp)

	if err != nil {
		fmt.Println("Failed to get stoct transfer")
		return resp, err
	}

	return resp, nil
}
