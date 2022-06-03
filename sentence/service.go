package sentence

import (
	"context"
	"errors"
)

var (
	ErrInvalidThird    = errors.New("invalid third")
	ErrServiceNotFound = errors.New("service not found")
)

type Service interface {
	GetSentence(ctx context.Context) (string, error)
}

type service struct {
	thirdServices map[string]Service
}

func NewService(thirdSvcs map[string]Service) Service {
	svc := new(service)
	svc.thirdServices = thirdSvcs

	// 可檢查 service 是否已實作 Service 介面所有方法
	return svc
}

func (svc *service) GetSentence(ctx context.Context) (string, error) {
	third, ok := ctx.Value(Third).(string)
	if !ok {
		return "", ErrInvalidThird
	}

	thirdSvc, ok := svc.thirdServices[third]
	if !ok {
		return "", ErrServiceNotFound
	}

	return thirdSvc.GetSentence(ctx)
}

type ServiceMiddleware func(Service) Service
