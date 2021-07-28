package espay

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
