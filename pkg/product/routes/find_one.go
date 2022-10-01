package routes

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jamalkaksouri/go-grpc-api-gateway/pkg/product/pb"
	"net/http"
	"strconv"
)

func FineOne(ctx *gin.Context, c pb.ProductServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	if err != nil {
		err := ctx.AbortWithError(http.StatusBadGateway, err)
		if err != nil {
			fmt.Println(http.ErrAbortHandler, err)
		}
		return
	}

	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: int64(id),
	})

	if err != nil {
		err := ctx.AbortWithError(http.StatusBadGateway, err)
		if err != nil {
			fmt.Println(http.ErrAbortHandler, err)
		}
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
