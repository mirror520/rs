package sentence

import "context"

type Service interface {
	GetSentence(ctx context.Context) (string, error)
}

type service struct {
}

func NewService() Service {
	return new(service) // 檢查 service 是否已實作 Service 介面所有方法
}

func (svc *service) GetSentence(ctx context.Context) (string, error) {
	return "", nil
}
