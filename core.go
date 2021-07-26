package sangu_espay

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
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
func (gateway *CoreGateway) Call(method, path string, header map[string]string, body io.Reader) ([]byte, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = gateway.Client.BaseUrl + path

	return gateway.Client.Call(method, path, header, body)
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
	req.Signature = gateway.GenerateSignature(req.RequuestUUID, req.RequestDateTime, req.OrderId, req.Amount, req.Ccy, req.MerchantCode)
	body := structToMap(&req)
	method := "POST"
	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
	}

	var responseBody []byte
	responseBody, err = gateway.Call(method, VA_PATH, headers, strings.NewReader(body.Encode()))
	if err != nil {
		gateway.Client.Logger.Error("Error during gateway.Call. Error : %v ", err)
		return
	}

	err = json.Unmarshal(responseBody, &res)
	if err != nil || res.ErrorCode != "00" {
		gateway.Client.Logger.Error("Error response error code %s is not equal to 00 or common error occurred : %v ", res.ErrorCode, err)
		return CreateVaResponse{}, errors.New(res.ErrorMessage)
	}

	return
}

func (gateway *CoreGateway) GenerateSignature(requuestUUID string, requestDateTime string, orderId string, amount string, currency string, merchantCode string) (signatureAsString string) {
	signature := "##" + gateway.Client.SignatureKey + "##" + requuestUUID +  "##" + requestDateTime +  "##" + orderId +  "##" + amount +  "##" + currency +  "##" + merchantCode +  "##" + SIGNATURE_MODE_SEND_INVOICE +  "##"
	signatureUpperCase := strings.ToUpper(signature)
	hash := sha256.Sum256([]byte(signatureUpperCase))
	return fmt.Sprintf("%x", hash[:])
}

func (gateway *CoreGateway) SendInquiryResponse(res InquiryResponse) (err error) {
	method := "POST"
	body, err := json.Marshal(res)

	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
	}

	var responseBody []byte
	responseBody, err = gateway.Call(method, VA_PATH, headers, strings.NewReader(string(body)))

	if err != nil {
		return err
	}

	err = json.Unmarshal(responseBody, &res)
	if err != nil {
		return errors.New(res.ErrorMessage)
	}

	return
}