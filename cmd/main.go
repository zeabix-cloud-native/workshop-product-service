package main

import (
	"github.com/zeabix-cloud-native/workshop-product-service/internal/adapters/handlers/httpv1"
	"github.com/zeabix-cloud-native/workshop-product-service/internal/adapters/repository"
	"github.com/zeabix-cloud-native/workshop-product-service/internal/core/services"

	"github.com/zeabix-cloud-native/workshop-product-service/utils"

	"fmt"
	"log"
	"time"
)

func main() {
	// Config
	port := utils.GetEnv("PORT", "3002")
	delay := utils.GetEnv("DELAY", "2s")

	u, _ := time.ParseDuration(delay)

	log.Printf("Starting server at port %s\n", port)
	log.Printf("Delay: %s\n", delay)

	repo := repository.NewMapDBRepository()

	// Initial sample data
	utils.InitDB("sample.json", repo)

	s := services.NewProductService(repo, u)

	handler := httpv1.NewHandler(s)
	handler.Serve(fmt.Sprintf(":%s", port))
}
