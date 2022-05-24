package base

import (
	"context"

	"github.com/mirror520/rs/sentence"
)

type service struct {
	thirdService sentence.Service
}

func NewService(third sentence.Service) sentence.Service {
	svc := new(service)
	svc.thirdService = third

	// 可檢查 service 是否已實作 Service 介面所有方法
	return svc
}

func (svc *service) GetSentence(ctx context.Context) (string, error) {
	return svc.thirdService.GetSentence(ctx)
}
