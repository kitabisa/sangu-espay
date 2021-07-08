package sangu_espay

import (
	"encoding/json"
	"io"
	"strings"
)

const (
	VA_PATH      = "rest/merchantpg/sendinvoice"
)

// CoreGateway struct
type CoreGateway struct {
	Client Client
}

// Call : base method to call Espay
func (gateway *CoreGateway) Call(method, path string, header map[string]string, body io.Reader, v interface{}, vErr interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = gateway.Client.BaseUrl + path

	return gateway.Client.Call(method, path, header, body, v, vErr)
}

func (gateway *CoreGateway) CreateVA(req CreateVaRequest) (res CreateVaResponse, err error) {
	method := "POST"
	body, err := json.Marshal(req)

	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
	}

	err = gateway.Call(method, VA_PATH, headers, strings.NewReader(string(body)), &res, nil)

	if err != nil {
		return
	}

	return
}

func (gateway *CoreGateway) SendInquiryResponse(inquiryRequest InquiryRequest) (res InquiryResponse, err error) {
	method := "POST"
	body, err := json.Marshal(inquiryRequest)

	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
	}

	err = gateway.Call(method, VA_PATH, headers, strings.NewReader(string(body)), &res, nil)

	if err != nil {
		return
	}

	return
}