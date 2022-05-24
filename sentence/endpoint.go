package sentence

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type ContextKey int

const Third ContextKey = iota

func GetSentenceEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.GetSentence(ctx)
	}
}
