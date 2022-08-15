package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/petersonsalme/go-grpc-project/go-grpc-api-gateway/internal/config"
	"github.com/petersonsalme/go-grpc-project/go-grpc-api-gateway/internal/routes/auth"
	"github.com/petersonsalme/go-grpc-project/go-grpc-api-gateway/internal/routes/order"
	"github.com/petersonsalme/go-grpc-project/go-grpc-api-gateway/internal/routes/product"
)

func RegisterAll(r *gin.Engine, c *config.Config) {
	auth := auth.RegisterRoutes(r, c)
	product.RegisterRoutes(r, c, auth)
	order.RegisterRoutes(r, c, auth)
}
