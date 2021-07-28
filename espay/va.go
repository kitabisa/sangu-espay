package espay

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

const (
	VaPath                   = "rest/merchantpg/sendinvoice"
	SignatureModeSendInvoice = "SENDINVOICE"
)

// createVaRequestBody Modify CreateVaRequest when  you change this struct
func createVaRequestBody(req CreateVaRequest) (values url.Values) {
	values = url.Values{}
	values.Set("rq_uuid", req.RequestUUID)
	values.Set("rq_datetime", req.RequestDateTime)
	values.Set("order_id", req.OrderId)
	values.Set("amount", req.Amount)
	values.Set("ccy", req.Ccy)
	values.Set("comm_code", req.MerchantCode)
	values.Set("remark1", req.Remark1)
	values.Set("remark2", req.Remark2)
	values.Set("remark3", req.Remark3)
	values.Set("update", req.Update)
	values.Set("bank_code", req.BankCode)
	values.Set("va_expired", req.VaExpired)
	values.Set("signature", req.Signature)
	return
}

// CreateVaRequest Modify createVaRequestBody when  you change this struct
type CreateVaRequest struct {
	RequestUUID     string `json:"rq_uuid" valid:"required"`
	RequestDateTime string `json:"rq_datetime" valid:"required"`
	OrderId         string `json:"order_id" valid:"required"`
	Amount          string `json:"amount" valid:"required"`
	Ccy             string `json:"ccy" valid:"required"`
	MerchantCode    string `json:"comm_code" valid:"required"`
	Remark1         string `json:"remark1"`
	Remark2         string `json:"remark2" valid:"required"`
	Remark3         string `json:"remark3"`
	Update          string `json:"update" valid:"required"`
	BankCode        string `json:"bank_code" valid:"required"`
	VaExpired       string `json:"va_expired"`
	Signature       string `json:"signature" valid:"required"`
}

type CreateVaResponse struct {
	RequestUUID     string `json:"rq_uuid" valid:"required"`
	RequestDateTime string `json:"rq_datetime" valid:"required"`
	ErrorCode       string `json:"error_code" valid:"required"`
	ErrorMessage    string `json:"error_message" valid:"required"`
	VaNumber        string `json:"va_number"`
	Expired         string `json:"expired"`
	Description     string `json:"desc"`
	TotalAmount     string `json:"total_amt"`
	BankCode        string `json:"bank_code"`
	Amount          string `json:"amount"`
	Fee             string `json:"fee"`
}

func (c *EspayClient) CreateVA(req CreateVaRequest) (res CreateVaResponse, err error) {
	body := createVaRequestBody(req)
	method := "POST"
	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
	}

	var responseBody []byte
	responseBody, err = c.Call(method, VaPath, headers, strings.NewReader(body.Encode()))
	if err != nil {
		c.Logger.Error("Error during gateway.Call. Error : %v ", err)
		return
	}

	err = json.Unmarshal(responseBody, &res)
	if err != nil || res.ErrorCode != "00" {
		c.Logger.Error("Error response error code %s is not equal to 00 or common error occurred : %v ", res.ErrorCode, err)
		return CreateVaResponse{}, errors.New(res.ErrorMessage)
	}

	return
}

func (vaRequest CreateVaRequest) CreateSignature(signatureKey string) string{
	signature := "##" + signatureKey + "##" + vaRequest.RequestUUID +  "##" + vaRequest.RequestDateTime +  "##" + vaRequest.OrderId +  "##" + vaRequest.Amount +  "##" + vaRequest.Ccy +  "##" + vaRequest.MerchantCode +  "##" + SignatureModeSendInvoice +  "##"
	signatureUpperCase := strings.ToUpper(signature)
	hash := sha256.Sum256([]byte(signatureUpperCase))
	return fmt.Sprintf("%x", hash[:])
}
