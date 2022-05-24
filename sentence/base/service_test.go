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

	thirdServices := make(map[string]sentence.Service)
	thirdServices["metaphorpsum"] = thirdSvc

	ctx := context.WithValue(context.Background(), sentence.Third, "metaphorpsum")

	svc := NewService(thirdServices)
	sentence, _ := svc.GetSentence(ctx)

	fmt.Println(sentence)
	assert.NotEmpty(sentence)
}

func TestSentenceServiceInjectItsthisforthatProxy(t *testing.T) {
	assert := assert.New(t)

	var thirdSvc sentence.Service // dummy service
	thirdSvc = metaphorpsum.ProxyMiddleware()(thirdSvc)

	thirdServices := make(map[string]sentence.Service)
	thirdServices["itsthisforthat"] = thirdSvc

	ctx := context.WithValue(context.Background(), sentence.Third, "itsthisforthat")

	svc := NewService(thirdServices)
	sentence, _ := svc.GetSentence(ctx)

	fmt.Println(sentence)
	assert.NotEmpty(sentence)
}
