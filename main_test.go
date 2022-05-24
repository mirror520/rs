package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/mirror520/rs/sentence"
	"github.com/mirror520/rs/sentence/third/metaphorpsum"

	"github.com/stretchr/testify/assert"
)

// TODO: 應該在 sentence/service_test.go 去測試
func TestSentenceService(t *testing.T) {
	assert := assert.New(t)

	var thirdSvc sentence.Service // dummy service
	thirdSvc = metaphorpsum.ProxyMiddleware()(thirdSvc)

	svc := sentence.NewService(thirdSvc)
	sentence, _ := svc.GetSentence(context.TODO())

	fmt.Println(sentence)
	assert.NotEmpty(sentence)
}
