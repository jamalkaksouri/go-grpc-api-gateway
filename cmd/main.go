package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jamalkaksouri/go-grpc-api-gateway/pkg/auth"
	"github.com/jamalkaksouri/go-grpc-api-gateway/pkg/config"
	"github.com/jamalkaksouri/go-grpc-api-gateway/pkg/order"
	"github.com/jamalkaksouri/go-grpc-api-gateway/pkg/product"
	"log"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	
	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	err = r.Run(c.Port)
	if err != nil {
		log.Println("api gateway run failed!")
	}
}
