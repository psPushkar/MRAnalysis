package main

import (
	demo "demo/internal"
	"html/template"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.SetFuncMap(template.FuncMap{
		"upper": strings.ToUpper,
	})

	//router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	router.POST("api/keyword", demo.GetAnalysisForKeyword)
	router.GET("/", demo.GetKeywords)

	if err := router.Run(":7777"); err != nil {
		log.Fatal("unable to start the server ")
	}

}
