package services

import (
	"github.com/zeabix-cloud-native/workshop-product-service/internal/core/domain"
	"github.com/zeabix-cloud-native/workshop-product-service/internal/core/ports"

	"time"
)

type service struct {
	r     ports.ProductRepository
	delay time.Duration
}

func NewProductService(repo ports.ProductRepository, delay time.Duration) ports.ProductService {
	return &service{
		r:     repo,
		delay: delay,
	}
}

func (s *service) ListProduct() ([]*domain.Product, error) {
	time.Sleep(s.delay)
	return s.r.ListProduct(), nil
}

func (s *service) GetProductDetail(id string) (*domain.Product, error) {
	time.Sleep(s.delay)
	return s.r.GetProductById(id)
}
