package sentence

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
)

func GetSentenceHandler(endpoint endpoint.Endpoint) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response, err := endpoint(ctx, nil)
		if err != nil {
			ctx.AbortWithError(http.StatusServiceUnavailable, err)
			return
		}

		resp := response.(string)
		ctx.String(http.StatusOK, resp)
	}
}
