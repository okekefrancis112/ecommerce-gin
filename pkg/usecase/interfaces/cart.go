package interfaces

import (
	"context"

	"github.com/okekefrancis112/ecommerce-gin/pkg/api/handler/request"
	"github.com/okekefrancis112/ecommerce-gin/pkg/api/handler/response"
	"github.com/okekefrancis112/ecommerce-gin/pkg/domain"
)

type CartUseCase interface {
	SaveProductItemToCart(ctx context.Context, userID, productItemId uint) error         // save product_item to cart
	RemoveProductItemFromCartItem(ctx context.Context, userID, productItemId uint) error // remove product_item from cart
	UpdateCartItem(ctx context.Context, updateDetails request.UpdateCartItem) error      // edit cartItems( quantity change )
	GetUserCart(ctx context.Context, userID uint) (cart domain.Cart, err error)
	GetUserCartItems(ctx context.Context, cartId uint) (cartItems []response.CartItem, err error)
}
