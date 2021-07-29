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
