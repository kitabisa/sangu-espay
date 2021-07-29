package sangu_espay

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

const (
	SIGNATURE_MODE_SEND_INVOICE                  = "SENDINVOICE"
	SIGNATURE_MODE_PAYMENT_NOTIFICATION_REQUEST  = "PAYMENTREPORT"
	SIGNATURE_MODE_PAYMENT_NOTIFICATION_RESPONSE = "PAYMENTREPORT-RS"
	SIGNATURE_MODE_INQUIRY_TRANSACTION_REQUEST   = "INQUIRY"
	SIGNATURE_MODE_INQUIRY_TRANSACTION_RESPONSE  = "INQUIRY-RS"
)

func (gateway *CoreGateway) GenerateSignatureCreateVARequest(requestUUID string, requestDateTime string, orderId string, amount string, currency string, merchantCode string) (signatureAsString string) {
	signature := "##" + gateway.Client.SignatureKey + "##" + requestUUID + "##" + requestDateTime + "##" + orderId + "##" + amount + "##" + currency + "##" + merchantCode + "##" + SIGNATURE_MODE_SEND_INVOICE + "##"
	signatureUpperCase := strings.ToUpper(signature)
	hash := sha256.Sum256([]byte(signatureUpperCase))
	return fmt.Sprintf("%x", hash[:])
}

func (gateway *CoreGateway) GenerateSignaturePaymentNotificationRequest(requestDateTime string, orderID string) (signatureAsString string) {
	signature := "##" + gateway.Client.SignatureKey + "##" + requestDateTime + "##" + orderID + "##" + SIGNATURE_MODE_PAYMENT_NOTIFICATION_REQUEST + "##"
	signatureUpperCase := strings.ToUpper(signature)
	hash := sha256.Sum256([]byte(signatureUpperCase))
	return fmt.Sprintf("%x", hash[:])
}

func (gateway *CoreGateway) GenerateSignaturePaymentNotificationResponse(requestUUID string, requestDateTime string, errorCode string) (signatureAsString string) {
	signature := "##" + gateway.Client.SignatureKey + "##" + requestUUID + "##" + requestDateTime + "##" + errorCode + "##" + SIGNATURE_MODE_PAYMENT_NOTIFICATION_RESPONSE + "##"
	signatureUpperCase := strings.ToUpper(signature)
	hash := sha256.Sum256([]byte(signatureUpperCase))
	return fmt.Sprintf("%x", hash[:])
}

func (gateway *CoreGateway) GenerateSignatureInquiryTransactionRequest(requestDateTime string, orderID string) (signatureAsString string) {
	signature := "##" + gateway.Client.SignatureKey + "##" + requestDateTime + "##" + orderID + "##" + SIGNATURE_MODE_INQUIRY_TRANSACTION_REQUEST + "##"
	signatureUpperCase := strings.ToUpper(signature)
	hash := sha256.Sum256([]byte(signatureUpperCase))
	return fmt.Sprintf("%x", hash[:])
}

func (gateway *CoreGateway) GenerateSignatureInquiryTransactionResponse(requestUUID string, requestDateTime string, orderID string, errorCode string) (signatureAsString string) {
	signature := "##" + gateway.Client.SignatureKey + "##" + requestUUID + "##" + requestDateTime + "##" + orderID + "##" + errorCode + "##" + SIGNATURE_MODE_INQUIRY_TRANSACTION_RESPONSE + "##"
	signatureUpperCase := strings.ToUpper(signature)
	hash := sha256.Sum256([]byte(signatureUpperCase))
	return fmt.Sprintf("%x", hash[:])
}
