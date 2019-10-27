package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Anthony-Hoo/dl-siteBackEnd/src/api"
)

func main()  {
	router := gin.Default()
	router.LoadHTMLGlob("static/*.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.POST("/switch2admin", api.Switch2admin)
}