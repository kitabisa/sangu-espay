package sangu_espay

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"reflect"
	"strings"
)

const (
	VA_PATH      = "rest/merchantpg/sendinvoice"
	SIGNATURE_MODE_SEND_INVOICE = "SENDINVOICE"
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

func structToMap(i interface{}) (values url.Values) {
	values = url.Values{}
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		values.Set(typ.Field(i).Tag.Get("json"), fmt.Sprint(iVal.Field(i)))
	}
	return
}

func (gateway *CoreGateway) CreateVA(req CreateVaRequest) (res CreateVaResponse, err error) {
	signature := generateSignature(gateway.Client.SignatureKey, req)
	req.Signature = fmt.Sprintf("%x", signature)
	body := structToMap(&req)
	method := "POST"
	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
	}

	err = gateway.Call(method, VA_PATH, headers, strings.NewReader(body.Encode()), &res, nil)

	if res.ErrorCode == "00" && res.ErrorMessage == "Success" {
		return res, nil
	} 

	if err != nil {
		return
	}

	return
}

func generateSignature(signatureKey string, req CreateVaRequest) []byte {
	signature := "##" + signatureKey + "##" + req.RequuestUUID + "##" + req.RequestDateTime + "##" + req.OrderId + "##" + req.Amount + "##" + req.Ccy + "##" + req.MerchantCode + "##" + SIGNATURE_MODE_SEND_INVOICE + "##"
	signatureUpperCase := strings.ToUpper(signature)
	hash := sha256.Sum256([]byte(signatureUpperCase))
	return hash[:]
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