package espay

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

const (
	SIGNATURE_MODE_PAYMENT_NOTIFICATION_REQUEST  = "PAYMENTREPORT"
	SIGNATURE_MODE_PAYMENT_NOTIFICATION_RESPONSE = "PAYMENTREPORT-RS"
)

type PaymentNotificationRequest struct {
	RequestUUID                string `json:"rq_uuid" schema:"rq_uuid" valid:"required"`
	RequestDateTime            string `json:"rq_datetime" schema:"rq_datetime" valid:"required"`
	SenderID                   string `json:"sender_id" schema:"sender_id"`
	ReceiverID                 string `json:"receiver_id" schema:"receiver_id"`
	Password                   string `json:"password" schema:"password"`
	MerchantCode               string `json:"comm_code" schema:"comm_code" valid:"required"`
	MemberCode                 string `json:"member_code" schema:"member_code"`
	MemberCustID               string `json:"member_cust_id" schema:"member_cust_id"`
	MemberCustName             string `json:"member_cust_name" schema:"member_cust_name"`
	Ccy                        string `json:"ccy" schema:"ccy" valid:"required"`
	Amount                     string `json:"amount" schema:"amount" valid:"required"`
	DebitFromBank              string `json:"debit_from_bank" schema:"debit_from_bank" valid:"required"`
	DebitFrom                  string `json:"debit_from" schema:"debit_from"`
	DebitFromName              string `json:"debit_from_name" schema:"debit_from_name"`
	CreditToBank               string `json:"credit_to_bank" schema:"credit_to_bank" valid:"required"`
	CreditTo                   string `json:"credit_to" schema:"credit_to"`
	CreditToName               string `json:"credit_to_name" schema:"credit_to_name"`
	PaymentDateTime            string `json:"payment_datetime" schema:"payment_datetime" valid:"required"`
	PaymentRef                 string `json:"payment_ref" schema:"payment_ref" valid:"required"`
	PaymentRemark              string `json:"payment_remark" schema:"payment_remark"`
	OrderID                    string `json:"order_id" schema:"order_id" valid:"required"`
	ProductCode                string `json:"product_code" schema:"product_code" valid:"required"`
	ProductValue               string `json:"product_value" schema:"product_value"`
	Message                    string `json:"message" schema:"message"`
	Status                     string `json:"status" schema:"status"`
	Token                      string `json:"token" schema:"token"`
	TotalAmount                string `json:"total_amount" schema:"total_amount"`
	TxKey                      string `json:"tx_key" schema:"tx_key"`
	FeeType                    string `json:"fee_type" schema:"fee_type"`
	TxFee                      string `json:"tx_fee" schema:"tx_fee"`
	ApprovalCode               string `json:"approval_code" schema:"approval_code"`
	MemberID                   string `json:"member_id" schema:"member_id"`
	ApprovalCodeFullBca        string `json:"approval_code_full_bca" schema:"approval_code_full_bca"`
	ApprovalCodeInstallmentBca string `json:"approval_code_installment_bca" schema:"approval_code_installment_bca"`
	Signature                  string `json:"signature" schema:"signature" valid:"required"`
}

type PaymentNotificationResponse struct {
	RequestUUID       string `json:"rq_uuid" valid:"required"`
	ResponseDateTime  string `json:"rs_datetime" valid:"required"`
	ErrorCode         string `json:"error_code"`
	ErrorMessage      string `json:"error_message" valid:"required"`
	Signature         string `json:"signature"`
	OrderId           string `json:"order_id"`
	ReconcileID       string `json:"reconcile_id" valid:"required"`
	ReconcileDateTime string `json:"reconcile_datetime" valid:"required"`
}

func (request PaymentNotificationRequest) CreateSignature(signatureKey string) string {
	signature := "##" + signatureKey + "##" + request.RequestDateTime + "##" + request.OrderID + "##" + SIGNATURE_MODE_PAYMENT_NOTIFICATION_REQUEST + "##"
	signatureUpperCase := strings.ToUpper(signature)
	hash := sha256.Sum256([]byte(signatureUpperCase))
	return fmt.Sprintf("%x", hash[:])
}

func (request PaymentNotificationResponse) CreateSignature(signatureKey string) string {
	signature := "##" + signatureKey + "##" + request.RequestUUID + "##" + request.ResponseDateTime + "##" + request.ErrorCode + "##" + SIGNATURE_MODE_PAYMENT_NOTIFICATION_RESPONSE + "##"
	signatureUpperCase := strings.ToUpper(signature)
	hash := sha256.Sum256([]byte(signatureUpperCase))
	return fmt.Sprintf("%x", hash[:])
}
