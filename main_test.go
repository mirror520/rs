package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/mirror520/rs/sentence"
	"github.com/mirror520/rs/sentence/third/metaphorpsum"

	"github.com/stretchr/testify/suite"
)

type SentenceServiceTestSuite struct {
	suite.Suite
}

func (suite *SentenceServiceTestSuite) TestServiceInjectMetaphorpsumProxy() {
	var thirdSvc sentence.Service // dummy service
	thirdSvc = metaphorpsum.ProxyMiddleware()(thirdSvc)

	thirdServices := make(map[string]sentence.Service)
	thirdServices["metaphorpsum"] = thirdSvc

	ctx := context.WithValue(context.Background(), sentence.Third, "metaphorpsum")

	svc := sentence.NewService(thirdServices)
	sentence, _ := svc.GetSentence(ctx)

	fmt.Println(sentence)
	suite.NotEmpty(sentence)
}

func (suite *SentenceServiceTestSuite) TestInjectItsthisforthatProxy() {
	var thirdSvc sentence.Service // dummy service
	thirdSvc = metaphorpsum.ProxyMiddleware()(thirdSvc)

	thirdServices := make(map[string]sentence.Service)
	thirdServices["itsthisforthat"] = thirdSvc

	ctx := context.WithValue(context.Background(), sentence.Third, "itsthisforthat")

	svc := sentence.NewService(thirdServices)
	sentence, _ := svc.GetSentence(ctx)

	fmt.Println(sentence)
	suite.NotEmpty(sentence)
}

func TestRunSentenceServiceSuite(t *testing.T) {
	suite.Run(t, new(SentenceServiceTestSuite))
}
