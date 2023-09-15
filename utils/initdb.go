package utils

import (
	"github.com/zeabix-cloud-native/workshop-product-service/internal/core/domain"
	"github.com/zeabix-cloud-native/workshop-product-service/internal/core/ports"

	"encoding/json"
	"io/ioutil"
	"os"
)

type Products struct {
	Products []domain.Product `json:"products"`
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	SKU         string  `json:"sku"`
	Price       float32 `json:"price"`
	Image       string  `json:"image"`
}

func InitDB(f string, repo ports.ProductRepository) error {
	jsonFile, err := os.Open(f)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var products Products
	json.Unmarshal(byteValue, &products)

	for _, p := range products.Products {
		dp := new(domain.Product)
		dp.ID = p.ID
		dp.Name = p.Name
		dp.Description = p.Description
		dp.SKU = p.SKU
		dp.Price = p.Price
		dp.Image = p.Image

		repo.Save(dp)
	}

	return nil

}
