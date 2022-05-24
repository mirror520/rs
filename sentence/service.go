package sentence

import (
	"context"
)

type Service interface {
	GetSentence(ctx context.Context) (string, error)
}

type service struct {
	thirdService Service
}

func NewService(third Service) Service {
	svc := new(service)
	svc.thirdService = third

	// 可檢查 service 是否已實作 Service 介面所有方法
	return svc
}

func (svc *service) GetSentence(ctx context.Context) (string, error) {
	return svc.thirdService.GetSentence(ctx)
}

type ServiceMiddleware func(Service) Service
