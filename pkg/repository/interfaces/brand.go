package interfaces

import (
	"github.com/okekefrancis112/ecommerce-gin/pkg/api/handler/request"
	"github.com/okekefrancis112/ecommerce-gin/pkg/domain"
)

type BrandRepository interface {
	IsExist(brand domain.Brand) (bool, error)
	Save(brand domain.Brand) (domain.Brand, error)
	Update(brand domain.Brand) error
	FindAll(pagination request.Pagination) ([]domain.Brand, error)
	FindOne(brandID uint) (domain.Brand, error)
	Delete(brandID uint) error
}
