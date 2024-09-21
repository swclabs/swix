package vnpay

type request struct{}

type response struct{}

type IVNPay interface {
	PaymentUrl(vnpayPaymentUrl string, secretKey string) string
	ValidateResp(secretKey string) bool
}

var _ IVNPay = (*VNPay)(nil)

type VNPay struct {
	req  request
	resp response
}

// PaymentUrl implements IVNPay.
func (v *VNPay) PaymentUrl(vnpayPaymentUrl string, secretKey string) string {
	panic("unimplemented")
}

// ValidateResp implements IVNPay.
func (v *VNPay) ValidateResp(secretKey string) bool {
	panic("unimplemented")
}

func HmacSHA512(key string, data []byte) string {
	return ""
}
