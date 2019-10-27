package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	
)

func main()  {
	router := gin.Default()
	router.LoadHTMLGlob("static/*.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.POST("/switch2admin",
}