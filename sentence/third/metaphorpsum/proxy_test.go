package metaphorpsum

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mirror520/rs/sentence"
)

func TestGetSentence(t *testing.T) {
	assert := assert.New(t)

	var svc sentence.Service     // dummy service
	svc = ProxyMiddleware()(svc) // ProxyMiddleware 實作相同介面，故亦可視為是 sentence.Service
	sentence, _ := svc.GetSentence(context.TODO())

	fmt.Println(sentence)
	assert.NotEmpty(sentence)
}
