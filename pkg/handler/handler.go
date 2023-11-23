package handler

import (
	_ "github.com/GRRashid/lamoda/docs"
	"github.com/GRRashid/lamoda/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		product := api.Group("/products")
		{
			product.POST("/create", h.createProduct)
			product.PUT("/reserve", h.reserveProduct)
			product.PUT("/unreserved", h.unreservedProduct)
		}

		storage := api.Group("/storages")
		{
			storage.POST("/create", h.createStorage)
			storage.GET(":storageId/products/unreserved", h.getAvailableProducts)
		}
	}

	return router
}
