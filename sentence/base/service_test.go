package base

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mirror520/rs/sentence"
	"github.com/mirror520/rs/sentence/third/metaphorpsum"
)

func TestSentenceServiceInjectMetaphorpsumProxy(t *testing.T) {
	assert := assert.New(t)

	var thirdSvc sentence.Service // dummy service
	thirdSvc = metaphorpsum.ProxyMiddleware()(thirdSvc)

	svc := NewService(thirdSvc)
	sentence, _ := svc.GetSentence(context.TODO())

	fmt.Println(sentence)
	assert.NotEmpty(sentence)
}
