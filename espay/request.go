package espay


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
	Password                   string `json:"password" schema:"password"`
	Signature                  string `json:"signature" schema:"signature" valid:"required"`
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
