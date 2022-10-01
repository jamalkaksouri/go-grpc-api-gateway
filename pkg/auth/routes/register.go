package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jamalkaksouri/go-grpc-api-gateway/pkg/auth/pb"
	"net/http"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	b := RegisterRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:    b.Email,
		Password: b.Password,
	})

	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
