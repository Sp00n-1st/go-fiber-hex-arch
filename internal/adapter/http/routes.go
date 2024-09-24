package http

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-hex-arch/internal/adapter/config"
	"go-fiber-hex-arch/internal/core/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(app *fiber.App, productService *service.ProductService, mongoClient *mongo.Client, cfg *config.HTTP) {
	productHandler := NewProductHandler(*productService)
	monitoringHandler := NewMonitoringHandler(mongoClient)

	app.Post(cfg.Prefix+"/products", productHandler.CreateProduct)
	app.Put(cfg.Prefix+"/products/:id", productHandler.UpdateProduct)
	app.Delete(cfg.Prefix+"/products/:id", productHandler.DeleteProduct)
	app.Get(cfg.Prefix+"/products/:id", productHandler.GetProductByID)
	app.Get(cfg.Prefix+"/products", productHandler.GetProducts)

	app.Get(cfg.Prefix+"/monitoring", monitoringHandler.GetMonitoringData)
}
