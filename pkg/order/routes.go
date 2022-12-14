package order

import (
	"github.com/gin-gonic/gin"
	"github.com/jamalkaksouri/go-grpc-api-gateway/pkg/auth"
	"github.com/jamalkaksouri/go-grpc-api-gateway/pkg/config"
	"github.com/jamalkaksouri/go-grpc-api-gateway/pkg/order/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routesGroup := r.Group("/order")
	routesGroup.Use(a.AuthRequired)
	routesGroup.POST("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
