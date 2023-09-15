package repository

import (
	"github.com/zeabix-cloud-native/workshop-product-service/internal/core/domain"
	"github.com/zeabix-cloud-native/workshop-product-service/internal/core/ports"
)

type mapdb struct {
	DB map[string]*domain.Product
}

func NewMapDBRepository() ports.ProductRepository {
	return &mapdb{
		DB: make(map[string]*domain.Product),
	}
}

func (r *mapdb) Save(product *domain.Product) error {
	p, ok := r.DB[product.ID]

	if !ok {
		// New item
		r.DB[product.ID] = product
		return nil
	}

	// Update existing item
	p.Name = product.Name
	p.Description = product.Description
	p.SKU = product.SKU
	p.Price = product.Price
	p.Image = product.Image

	return nil
}

func (r *mapdb) GetProductById(id string) (*domain.Product, error) {
	p, ok := r.DB[id]
	if !ok {
		return nil, ports.ErrProductNotFound
	}

	return p, nil
}

func (r *mapdb) ListProduct() []*domain.Product {
	products := make([]*domain.Product, 0, len(r.DB))
	for _, v := range r.DB {
		products = append(products, v)
	}

	return products
}
