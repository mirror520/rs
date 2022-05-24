package metaphorpsum

import (
	"context"

	"github.com/mirror520/rs/sentence"
)

func ProxyMiddleware() sentence.ServiceMiddleware {
	return func(next sentence.Service) sentence.Service {
		return new(proxyMiddleware)
	}
}

type proxyMiddleware struct {
	next sentence.Service
}

func (mw *proxyMiddleware) GetSentence(ctx context.Context) (string, error) {
	return "", nil
}
