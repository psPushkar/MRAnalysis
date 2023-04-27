package main

import (
	demo "demo/internal"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("api/keyword", demo.GetAnalysisForKeyword)

	if err := router.Run(":7777"); err != nil {
		log.Fatal("unable to start the server ")
	}

}
