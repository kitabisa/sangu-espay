package espay

import (
	"crypto/sha256"
	"fmt"
	"log"
	"strings"
)

type InquiryRequest struct {
	RequestUUID     string `json:"rq_uuid" valid:"required"`
	RequestDateTime string `json:"rq_datetime" valid:"required"`
	MemberId        string `json:"member_id" valid:"-"`
	MerchantCode    string `json:"comm_code" valid:"required"`
	OrderId         string `json:"order_id" valid:"required"`
	Password        string `json:"password" valid:"-"`
	Signature       string `json:"signature" valid:"required"`
}

type InquiryResponse struct {
	RequestUUID       string `json:"rq_uuid" valid:"required"`
	ResponseDateTime   string `json:"rs_datetime" valid:"required"`
	ErrorCode         string `json:"error_code" valid:"required"`
	ErrorMessage      string `json:"error_message" valid:"required"`
	Signature         string `json:"signature" valid:"required"`
	OrderId           string `json:"order_id" valid:"-"`
	Amount            string `json:"amount" valid:"-"`
	Ccy               string `json:"ccy" valid:"-"`
	Description       string `json:"desc" valid:"-"`
	TransactionDate   string `json:"trx_date" valid:"-"`
	InstallmentPeriod string `json:"installment_period" valid:"-"`
	CustomerDetails   CustomerDetails
	ShippingAddress   ShippingAddress
}

type CustomerDetails struct {
	FirstName string `json:"firstname" valid:"required"`
	LastName  string `json:"lastname" valid:"-"`
	Phone     string `json:"phone_number" valid:"required"`
	Email     string `json:"email" valid:"required"`
}

type ShippingAddress struct {
	FirstName   string `json:"firstname" valid:"required"`
	LastName    string `json:"lastname" valid:"-"`
	Address     string `json:"address" valid:"required"`
	City        string `json:"city" valid:"required"`
	PostalCode  string `json:"postal_code" valid:"required"`
	Phone       string `json:"phone_number" valid:"required"`
	CountryCode string `json:"country_code"  valid:"required"`
}

const SignatureModeInquiryTransactionRequest = "INQUIRY"
const SignatureModeInquiryTransactionResponse = "INQUIRY-RS"

func (inquiryResp InquiryResponse) CreateSignature(signatureKey string) string{
	signature := "##" + signatureKey + "##" + inquiryResp.RequestUUID + "##" + inquiryResp.ResponseDateTime + "##" +
		inquiryResp.OrderId + "##" + inquiryResp.ErrorCode + "##" + SignatureModeInquiryTransactionResponse + "##"
	signatureUpperCase := strings.ToUpper(signature)

	log.Println("InquiryResponse Signature Format", signatureUpperCase)
	hash := sha256.Sum256([]byte(signatureUpperCase))
	return fmt.Sprintf("%x", hash[:])
}

func (inquiryReq InquiryRequest) CreateSignature(signatureKey string) (signatureAsString string) {
	signature := "##" + signatureKey  + "##" + inquiryReq.RequestDateTime + "##" + inquiryReq.OrderId + "##" + SignatureModeInquiryTransactionRequest + "##"
	signatureUpperCase := strings.ToUpper(signature)

	log.Println("InquiryRequest Signature Format", signatureUpperCase)
	hash := sha256.Sum256([]byte(signatureUpperCase))
	return fmt.Sprintf("%x", hash[:])
}
