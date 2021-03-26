package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router := gin.Default()
	initializeRoutes()
	router.LoadHTMLGlob("templates/*")
}

//func showIndexPage(c *gin.Context) {
//	articles := getAllArticles()
//
//	// Call the HTML method of the Context to render a template
//	c.HTML(
//		// Set the HTTP status to 200 (OK)
//		http.StatusOK,
//		// Use the index.html template
//		"index.html",
//		// Pass the data that the page uses
//		gin.H{
//			"title":   "Home Page",
//			"payload": articles,
//		},
//	)
//
//}