package httpv1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/zeabix-cloud-native/workshop-product-service/internal/core/ports"

	"encoding/json"
)

type ProductResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	SKU         string  `json:"sku"`
	Price       float32 `json:"price"`
	Image       string  `json:"image"`
}

type Handler struct {
	s   ports.ProductService
	app *fiber.App
}

func NewHandler(srv ports.ProductService) *Handler {
	return &Handler{
		s:   srv,
		app: fiber.New(),
	}
}

func (h *Handler) Serve(port string) error {
	h.app.Use(cors.New())
	v1 := h.app.Group("v1")
	v1.Get("/products", h.listProductHandler)
	v1.Get("/products/:id", h.getProductHandler)

	return h.app.Listen(port)
}

func (h *Handler) listProductHandler(c *fiber.Ctx) error {
	products, err := h.s.ListProduct()
	if err != nil {
		return fiber.ErrInternalServerError
	}

	res := make([]*ProductResponse, 0, len(products))

	for _, v := range products {
		p := new(ProductResponse)
		p.ID = v.ID
		p.Name = v.Name
		p.Description = v.Description
		p.SKU = v.SKU
		p.Price = v.Price
		p.Image = v.Image
		res = append(res, p)
	}

	resStr, err := json.Marshal(res)
	c.Status(fiber.StatusOK).SendString(string(resStr))
	return nil
}

func (h *Handler) getProductHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.ErrBadRequest
	}

	p, err := h.s.GetProductDetail(id)
	if err == ports.ErrProductNotFound {
		return fiber.ErrNotFound
	}

	if err != nil {
		return fiber.ErrInternalServerError
	}

	res := new(ProductResponse)
	res.ID = p.ID
	res.Name = p.Name
	res.Description = p.Description
	res.SKU = p.SKU
	res.Price = p.Price
	res.Image = p.Image

	resStr, err := json.Marshal(res)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	c.Status(fiber.StatusOK).SendString(string(resStr))
	return nil
}
