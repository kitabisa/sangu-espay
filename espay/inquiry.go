package espay

import (
	"crypto/sha256"
	"fmt"
	"net/url"
	"strings"
)

// createInquiryResponseBody Modify InquiryResponse when  you change this struct
func createInquiryResponseBody(req InquiryResponse) (values url.Values) {
	values = url.Values{}
	values.Set("rq_uuid", req.RequestUUID)
	values.Set("rq_datetime", req.RequestDateTime)
	values.Set("error_code", req.ErrorCode)
	values.Set("error_message", req.ErrorMessage)
	values.Set("signature", req.Signature)
	values.Set("order_id", req.OrderId)
	values.Set("amount", req.Amount)
	values.Set("ccy", req.Ccy)
	values.Set("description", req.Description)
	values.Set("trx_date", req.TransactionDate)
	values.Set("installment_period", req.InstallmentPeriod)
	values.Set("customer_details.firstname", req.CustomerDetails.FirstName)
	values.Set("customer_details.lastname", req.CustomerDetails.LastName)
	values.Set("customer_details.phone_number", req.CustomerDetails.Phone)
	values.Set("customer_details.email", req.CustomerDetails.Email)
	values.Set("shipping_address.first_name", req.ShippingAddress.FirstName)
	values.Set("shipping_address.lastname", req.ShippingAddress.LastName)
	values.Set("shipping_address.address", req.ShippingAddress.Address)
	values.Set("shipping_address.city", req.ShippingAddress.City)
	values.Set("shipping_address.postal_code", req.ShippingAddress.PostalCode)
	values.Set("shipping_address.phone", req.ShippingAddress.Phone)
	values.Set("shipping_address.country_code", req.ShippingAddress.CountryCode)
	return
}

// InquiryResponse Modify createInquiryResponseBody when  you change this struct
type InquiryResponse struct {
	RequestUUID       string `json:"rq_uuid" valid:"required"`
	RequestDateTime   string `json:"rq_datetime" valid:"required"`
	ErrorCode         string `json:"error_code" valid:"required"`
	ErrorMessage      string `json:"error_message" valid:"required"`
	Signature         string `json:"signature" valid:"required"`
	OrderId           string `json:"order_id"`
	Amount            string `json:"amount"`
	Ccy               string `json:"ccy" `
	Description       string `json:"desc"`
	TransactionDate   string `json:"trx_date"`
	InstallmentPeriod string `json:"installment_period"`
	CustomerDetails   CustomerDetails
	ShippingAddress   ShippingAddress
}

type CustomerDetails struct {
	FirstName string `json:"firstname" valid:"required"`
	LastName  string `json:"lastname"`
	Phone     string `json:"phone_number" valid:"required"`
	Email     string `json:"email" valid:"required"`
}

type ShippingAddress struct {
	FirstName   string `json:"firstname" valid:"required"`
	LastName    string `json:"lastname"`
	Address     string `json:"address" valid:"required"`
	City        string `json:"city" valid:"required"`
	PostalCode  string `json:"postal_code" valid:"required"`
	Phone       string `json:"phone_number" valid:"required"`
	CountryCode string `json:"country_code"  valid:"required"`
}


func (vaRequest InquiryResponse) CreateSignature(signatureKey string) string{
	signature := "##" + signatureKey + "##" + vaRequest.RequestUUID
	signatureUpperCase := strings.ToUpper(signature)
	hash := sha256.Sum256([]byte(signatureUpperCase))
	return fmt.Sprintf("%x", hash[:])
}
