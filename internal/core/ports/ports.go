package ports

import (
	"github.com/zeabix-cloud-native/workshop-product-service/internal/core/domain"

	"errors"
)

var (
	ErrProductNotFound = errors.New("product not found")
)

type ProductService interface {
	ListProduct() ([]*domain.Product, error)
	GetProductDetail(id string) (*domain.Product, error)
}

type ProductRepository interface {
	Save(product *domain.Product) error
	GetProductById(id string) (*domain.Product, error)
	ListProduct() []*domain.Product
}
