package main

import (
	"github.com/gin-gonic/gin"

	"github.com/mirror520/rs/sentence"
	"github.com/mirror520/rs/sentence/third/metaphorpsum"
)

func main() {
	var thirdSvc sentence.Service // dummy service
	thirdSvc = metaphorpsum.ProxyMiddleware()(thirdSvc)

	svc := sentence.NewService(thirdSvc)
	endpoint := sentence.GetSentenceEndpoint(svc)
	handler := sentence.GetSentenceHandler(endpoint)

	router := gin.Default()
	router.GET("/sentence", handler)
	router.Run(":8080")
}
