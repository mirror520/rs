package main

import (
	"github.com/gin-gonic/gin"

	"github.com/mirror520/rs/sentence"
	"github.com/mirror520/rs/sentence/base"
	"github.com/mirror520/rs/sentence/third/itsthisforthat"
	"github.com/mirror520/rs/sentence/third/metaphorpsum"
)

func main() {
	thirdServices := make(map[string]sentence.Service)

	var thirdSvc sentence.Service // dummy service
	thirdServices["metaphorpsum"] = metaphorpsum.ProxyMiddleware()(thirdSvc)
	thirdServices["itsthisforthat"] = itsthisforthat.ProxyMiddleware()(thirdSvc)

	var svc sentence.Service = base.NewService(thirdServices)
	endpoint := sentence.GetSentenceEndpoint(svc)
	handler := sentence.GetSentenceHandler(endpoint)

	router := gin.Default()
	router.GET("/:third/sentence", handler)
	router.Run(":8080")
}
