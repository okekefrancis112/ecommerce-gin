package response

import "github.com/okekefrancis112/ecommerce-gin/pkg/domain"

type OrderPayment struct {
	PaymentType  domain.PaymentType `json:"payment_type"`
	PaymentOrder any                `json:"payment_order"`
}
