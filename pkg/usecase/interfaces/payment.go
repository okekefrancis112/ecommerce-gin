package interfaces

import (
	"context"

	"github.com/okekefrancis112/ecommerce-gin/pkg/api/handler/request"
	"github.com/okekefrancis112/ecommerce-gin/pkg/api/handler/response"
	"github.com/okekefrancis112/ecommerce-gin/pkg/domain"
)

type PaymentUseCase interface {
	FindAllPaymentMethods(ctx context.Context) ([]domain.PaymentMethod, error)
	FindPaymentMethodByID(ctx context.Context, paymentMethodID uint) (domain.PaymentMethod, error)
	UpdatePaymentMethod(ctx context.Context, paymentMethod request.PaymentMethodUpdate) error

	// razorpay
	MakeRazorpayOrder(ctx context.Context, userID, shopOrderID uint) (razorpayOrder response.RazorpayOrder, err error)
	VerifyRazorPay(ctx context.Context, verifyReq request.RazorpayVerify) error
	// stipe
	MakeStripeOrder(ctx context.Context, userID, shopOrderID uint) (stipeOrder response.StripeOrder, err error)
	VerifyStripOrder(ctx context.Context, stripePaymentID string) error

	ApproveShopOrderAndClearCart(ctx context.Context, userID uint, approveDetails request.ApproveOrder) error
}
