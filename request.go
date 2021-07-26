package sangu_espay

type CreateVaRequest struct {
	RequuestUUID    string `json:"rq_uuid" valid:"required"`
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
	RequestUUID     string `json:"rq_uuid" valid:"required"`
	RequestDateTime string `json:"rq_datetime" valid:"required"`
	MemberId        string `json:"member_id"`
	MerchantCode    string `json:"comm_code" valid:"required"`
	OrderId         string `json:"order_id" valid:"required"`
	Password        string `json:"password"`
	Signature       string `json:"signature" valid:"required"`
}

type PaymentNotificationRequest struct {
	RequestUUID                string `json:"rq_uuid" valid:"required"`
	RequestDateTime            string `json:"rq_datetime" valid:"required"`
	Password                   string `json:"password"`
	Signature                  string `json:"signature" valid:"required"`
	MemberId                   string `json:"member_id"`
	MerchantCode               string `json:"comm_code" valid:"required"`
	OrderId                    string `json:"order_id" valid:"required"`
	Ccy                        string `json:"ccy" valid:"required"`
	Amount                     string `json:"amount" valid:"required"`
	DebitFromBank              string `json:"debit_from_bank" valid:"required"`
	DebitFrom                  string `json:"debit_from"`
	DebitFromName              string `json:"debit_from_name"`
	CreditToBank               string `json:"credit_to_bank" valid:"required"`
	CreditTo                   string `json:"credit_to"`
	CreditToName               string `json:"credit_to_name"`
	ProductCode                string `json:"product_code" valid:"required"`
	Message                    string `json:"message"`
	PaymentDateTime            string `json:"payment_datetime" valid:"required"`
	PaymentRef                 string `json:"payment_ref" valid:"required"`
	ApprovalCodeFullBca        string `json:"approval_code_full_bca"`
	ApprovalCodeInstallmentBca string `json:"approval_code_installment_bca"`
}
