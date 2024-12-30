package interfaces

import (
	"context"

	"github.com/okekefrancis112/ecommerce-gin/pkg/api/handler/request"
	"github.com/okekefrancis112/ecommerce-gin/pkg/api/handler/response"
	"github.com/okekefrancis112/ecommerce-gin/pkg/domain"
)

type AdminRepository interface {
	FindAdminByEmail(ctx context.Context, email string) (domain.Admin, error)
	FindAdminByUserName(ctx context.Context, userName string) (domain.Admin, error)
	SaveAdmin(ctx context.Context, admin domain.Admin) error

	FindAllUser(ctx context.Context, pagination request.Pagination) (users []response.User, err error)

	CreateFullSalesReport(ctc context.Context, reqData request.SalesReport) (salesReport []response.SalesReport, err error)

	//stock side
	FindStockBySKU(ctx context.Context, sku string) (stock response.Stock, err error)
}
