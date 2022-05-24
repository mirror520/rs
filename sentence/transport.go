package sentence

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
)

func GetSentenceHandler(endpoint endpoint.Endpoint) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		third := ctx.Param("third")

		newCtx := context.WithValue(context.Background(), Third, third)
		response, err := endpoint(newCtx, nil)
		if err != nil {
			ctx.AbortWithError(http.StatusServiceUnavailable, err)
			return
		}

		resp := response.(string)
		ctx.String(http.StatusOK, resp)
	}
}
