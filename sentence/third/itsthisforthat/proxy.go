package itsthisforthat

import (
	"context"

	"github.com/go-resty/resty/v2"

	"github.com/mirror520/rs/sentence"
)

func ProxyMiddleware() sentence.ServiceMiddleware {
	return func(next sentence.Service) sentence.Service {
		return &proxyMiddleware{next}
	}
}

type proxyMiddleware struct {
	next sentence.Service
}

func (mw *proxyMiddleware) GetSentence(ctx context.Context) (string, error) {
	client := resty.New().
		SetBaseURL("https://itsthisforthat.com")

	var sentence string
	resp, err := client.R().
		SetResult(&sentence).
		Get("/api.php?text")

	if err != nil {
		return "", err
	}

	return string(resp.Body()), nil
}
