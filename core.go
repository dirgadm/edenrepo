package edenrepo

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
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

// GetAllStockTransfer ...
func (auth *Auth) GetAllStockTransfer(values url.Values) (r *http.Response, e error) {
	resp := &http.Response{}

	// check whether url have querystring or not
	var uri = baseURL + "/v1/inventory/stock-transfer/vendor"
	queryString, e := queryStringBuilder(values)
	if e == nil {
		uri += queryString
	}

	req, err := http.NewRequest("GET", uri, nil)
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
	return resp, nil
}

// GetStockTransferByID ...
func (auth *Auth) GetStockTransferByID(id string) (r *http.Response, e error) {
	resp := &http.Response{}

	url := baseURL + "/v1/inventory/stock-transfer/" + id + "/vendor"
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
	return resp, nil
}

// InsertStockTransfer ...
func (auth *Auth) InsertStockTransfer(strnf *StockTransferRequest) (r *http.Response, e error) {
	jsonReq, e := json.Marshal(strnf)

	u, _ := url.ParseRequestURI(baseURL)
	u.Path = "/v1/inventory/stock-transfer/vendor"
	urlStr := u.String()

	//client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, urlStr, bytes.NewBuffer(jsonReq))
	req.Header.Add("Authorization", bearer+auth.Credentials)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}
	//cr := StockTransferRequest{}

	client := &http.Client{}
	resp, e := client.Do(req)
	if err != nil {
		log.Fatalln(e.Error())
		return
	}

	return resp, nil
}

// InsertStockTransfer ...
func (auth *Auth) UpdateStockTransfer(strnf *StockTransferRequest, id string) (r *http.Response, e error) {
	jsonReq, e := json.Marshal(strnf)

	u, _ := url.ParseRequestURI(baseURL)
	u.Path = "/v1/inventory/stock-transfer/" + id + "/vendor"
	urlStr := u.String()

	//client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, urlStr, bytes.NewBuffer(jsonReq))
	req.Header.Add("Authorization", bearer+auth.Credentials)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}
	//cr := StockTransferRequest{}

	client := &http.Client{}
	resp, e := client.Do(req)
	if err != nil {
		log.Fatalln(e.Error())
		return
	}

	return resp, nil
}

// PopulateData ..
func PopulateData(strnf *StockTransferRequest) (v url.Values) {
	data := url.Values{}
	data.Set("recognition_date", strnf.RecognitionDate)
	data.Set("origin_id", strnf.OriginID)
	data.Set("destination_id", strnf.DestinationID)
	data.Set("destination_address", strnf.DestinationAddress)
	data.Set("eta_date", strnf.EtaDate)
	data.Set("eta_time", strnf.EtaTime)
	data.Set("note", strnf.Note)
	data.Set("total_cost", strconv.Itoa(int(strnf.TotalCost)))

	return data
}
