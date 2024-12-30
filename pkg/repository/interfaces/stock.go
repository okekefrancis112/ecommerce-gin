package interfaces

import (
	"context"

	"github.com/okekefrancis112/ecommerce-gin/pkg/api/handler/request"
	"github.com/okekefrancis112/ecommerce-gin/pkg/api/handler/response"
)

type StockRepository interface {
	FindAll(ctx context.Context, pagination request.Pagination) (stocks []response.Stock, err error)
	Update(ctx context.Context, updateValues request.UpdateStock) error
}
