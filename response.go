package sangu_espay

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

type InquiryResponse struct {
	RequestUUID       string  `json:"rq_uuid" valid:"required"`
	RequestDateTime   string  `json:"rq_datetime" valid:"required"`
	ErrorCode         string  `json:"error_code" valid:"required"`
	ErrorMessage      string  `json:"error_message" valid:"required"`
	Signature         string  `json:"signature" valid:"required"`
	OrderId           string  `json:"order_id"`
	Amount            float64 `json:"amount"`
	Ccy               string  `json:"ccy" `
	Description       string  `json:"desc"`
	TransactionDate   string  `json:"trx_date"`
	InstallmentPeriod string  `json:"installment_period"`
	CustomerDetails   CustomerDetails
	ShippingAddress   ShippingAddress
}

type PaymentNotificationResponse struct {
	RequestUUID       string `json:"rq_uuid" valid:"required"`
	ResponseDateTime  string `json:"rs_datetime" valid:"required"`
	ErrorCode         string `json:"error_code"`
	ErrorMessage      string `json:"error_message" valid:"required"`
	Signature         string `json:"signature" valid:"required"`
	OrderId           string `json:"order_id"`
	ReconcileID       string `json:"reconcile_id" valid:"required"`
	ReconcileDateTime string `json:"reconcile_datetime" valid:"required"`
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
