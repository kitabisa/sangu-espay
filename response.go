package sangu_espay

type CreateVaResponse struct {
	RequuestUUID string `json:"rq_uuid" valid:"required"`
	RequestDateTime         string `json:"rq_datetime" valid:"required"`
	ErrorCode        string `json:"error_cd" valid:"required"`
	ErrorMessage            string `json:"error_msg" valid:"required"`
	VaNumber          string `json:"va_number" valid:"required"`
	Expired     string `json:"expired" valid:"required"`
	Description     string `json:"desc" valid:"required"`
	TotalAmount     string `json:"total_amt" valid:"required"`
	Amount     string `json:"amount" valid:"required"`
	Fee     string `json:"fee" valid:"required"`
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