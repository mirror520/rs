package sentence

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func GetSentenceEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.GetSentence(ctx)
	}
}
