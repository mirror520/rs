package sentence

import (
	"context"
)

type Service interface {
	GetSentence(ctx context.Context) (string, error)
}

type ServiceMiddleware func(Service) Service
