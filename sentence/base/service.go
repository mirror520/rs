package base

import (
	"context"
	"errors"

	"github.com/mirror520/rs/sentence"
)

var (
	ErrInvalidThird    = errors.New("invalid third")
	ErrServiceNotFound = errors.New("service not found")
)

type service struct {
	thirdServices map[string]sentence.Service
}

func NewService(thirdSvcs map[string]sentence.Service) sentence.Service {
	svc := new(service)
	svc.thirdServices = thirdSvcs

	// 可檢查 service 是否已實作 Service 介面所有方法
	return svc
}

func (svc *service) GetSentence(ctx context.Context) (string, error) {
	third, ok := ctx.Value(sentence.Third).(string)
	if !ok {
		return "", ErrInvalidThird
	}

	thirdSvc, ok := svc.thirdServices[third]
	if !ok {
		return "", ErrServiceNotFound
	}

	return thirdSvc.GetSentence(ctx)
}
