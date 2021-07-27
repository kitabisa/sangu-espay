package sangu_espay

import "net/url"

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

type InquiryRequest struct {
	RequestUUID     string `schema:"rq_uuid, required"`
	RequestDateTime string `schema:"rq_datetime, required"`
	MemberId        string `schema:"member_id"`
	MerchantCode    string `schema:"comm_code, required"`
	OrderId         string `schema:"order_id, required"`
	Password        string `schema:"password"`
	Signature       string `schema:"signature, required"`
}

type PaymentNotificationRequest struct {
	RequestUUID                string `json:"rq_uuid" schema:"rq_uuid" valid:"required"`
	RequestDateTime            string `json:"rq_datetime" schema:"rq_datetime" valid:"required"`
	Password                   string `json:"Password" schema:"Password"`
	Signature                  string `json:"Signature" schema:"Signature" valid:"required"`
	MemberID                   string `json:"member_id" schema:"member_id"`
	MerchantCode               string `json:"comm_code" schema:"comm_code" valid:"required"`
	OrderID                    string `json:"order_id" schema:"order_id" valid:"required"`
	Ccy                        string `json:"ccy" schema:"ccy" valid:"required"`
	Amount                     string `json:"amount" schema:"amount" valid:"required"`
	DebitFromBank              string `json:"debit_from_bank" schema:"debit_from_bank" valid:"required"`
	DebitFrom                  string `json:"debit_from" schema:"debit_from"`
	DebitFromName              string `json:"debit_from_name" schema:"debit_from_name"`
	CreditToBank               string `json:"credit_to_bank" schema:"credit_to_bank" valid:"required"`
	CreditTo                   string `json:"credit_to" schema:"credit_to"`
	CreditToName               string `json:"credit_to_name" schema:"credit_to_name"`
	ProductCode                string `json:"product_code" schema:"product_code" valid:"required"`
	Message                    string `json:"message" schema:"message"`
	PaymentDateTime            string `json:"payment_datetime" schema:"payment_datetime" valid:"required"`
	PaymentRef                 string `json:"payment_ref" schema:"payment_ref" valid:"required"`
	ApprovalCodeFullBca        string `json:"approval_code_full_bca" schema:"approval_code_full_bca"`
	ApprovalCodeInstallmentBca string `json:"approval_code_installment_bca" schema:"approval_code_installment_bca"`
}
