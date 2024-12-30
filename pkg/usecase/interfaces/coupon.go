package interfaces

import (
	"context"

	"github.com/okekefrancis112/ecommerce-gin/pkg/api/handler/request"
	"github.com/okekefrancis112/ecommerce-gin/pkg/api/handler/response"
	"github.com/okekefrancis112/ecommerce-gin/pkg/domain"
)

type CouponUseCase interface {
	// coupon
	AddCoupon(ctx context.Context, coupon domain.Coupon) error
	GetAllCoupons(ctx context.Context, pagination request.Pagination) (coupons []domain.Coupon, err error)
	UpdateCoupon(ctx context.Context, coupon domain.Coupon) error

	//user side coupons
	GetCouponsForUser(ctx context.Context, userID uint, pagination request.Pagination) (coupons []response.UserCoupon, err error)

	GetCouponByCouponCode(ctx context.Context, couponCode string) (coupon domain.Coupon, err error)
	ApplyCouponToCart(ctx context.Context, userID uint, couponCode string) (discountPrice uint, err error)
}
