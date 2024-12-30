package interfaces

import (
	"context"

	"github.com/okekefrancis112/ecommerce-gin/pkg/api/handler/request"
	"github.com/okekefrancis112/ecommerce-gin/pkg/domain"
)

type PaymentRepository interface {
	FindPaymentMethodByID(ctx context.Context, paymentMethodID uint) (paymentMethods domain.PaymentMethod, err error)
	FindPaymentMethodByType(ctx context.Context, paymentType domain.PaymentType) (paymentMethod domain.PaymentMethod, err error)
	FindAllPaymentMethods(ctx context.Context) ([]domain.PaymentMethod, error)
	UpdatePaymentMethod(ctx context.Context, paymentMethod request.PaymentMethodUpdate) error
}
