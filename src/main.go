package main

import (
	"net/http"

	"github.com/Anthony-Hoo/dl-siteBackEnd/src/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("src/template/*.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.POST("/switch2admin", api.Switch2Admin)

	router.Run("127.0.0.1:8080")
}
