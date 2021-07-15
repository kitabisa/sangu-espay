package sangu_espay

type CreateVaResponse struct {
	RequestUUID string `json:"rq_uuid" valid:"required"`
	RequestDateTime         string `json:"rq_datetime" valid:"required"`
	ErrorCode        string `json:"error_cd" valid:"required"`
	ErrorMessage            string `json:"error_msg" valid:"required"`
	VaNumber          string `json:"va_number"`
	Expired     string `json:"expired"`
	Description     string `json:"desc"`
	TotalAmount     string `json:"total_amt"`
	BankCode 	string `json:"bank_code"`
	Amount     string `json:"amount"`
	Fee     string `json:"fee"`
}

type InquiryResponse struct {
	ErrorCode string `json:"error_code" valid:"required"`
	ErrorMessage         string `json:"error_msg" valid:"required"`
	OrderId     string `json:"order_id"`
	Amount     float64 `json:"amount"`
	Ccy     string `json:"ccy" `
	Description     string `json:"desc"`
	TransactionDate     string `json:"trx_date" `
}