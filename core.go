package edenrepo

import (
	"log"
	"net/http"
)

// CoreGateway ...
type Auth struct {
	Credentials string
}

// Call : Base Method to call Core API
//func (client *CoreGateway) Call(method, path string, body io.Reader, v interface{}) error {
//	credentials := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpveUBlZGVuZmFybS5pZCIsInBhc3N3b3JkIjoiZGV2MTIzIn0.tgiwi4Mwcj5l2oyrCKIpN6BE7dHCO18vqBAnfk1Q4hU"
//	var bearer = "Bearer " + credentials
//
//	req, err := http.NewRequest(method, path, nil)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	req.Header.Add("Authorization", bearer)
//	resp, err := client.Client.Do(req)
//	if err != nil {
//		log.Println("Error on response.\n[ERRO] -", err)
//	}
//
//	defer resp.Body.Close()
//
//	return err
//}

// GetStockTransferByID ...
func (auth *Auth) GetStockTransferByID(id string) (r http.Response, e error) {
	if auth.Credentials == "" {
		log.Fatalln("missing or malformed jwt")
		r.Status = "fail"
		return
	}
	resp := &http.Response{}

	url := baseURL + "/v1/inventory/stock-transfer/" + id + ""
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", bearer+auth.Credentials)
	client := &http.Client{}
	resp, e = client.Do(req)
	if err != nil {
		log.Fatalln(e.Error())
		return
	}

	defer resp.Body.Close()

	return *resp, nil
}
