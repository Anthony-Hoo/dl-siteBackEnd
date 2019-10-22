package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"
	"net/http"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.POST("/switch2admin", func(c *gin.Context) {
		passwd := c.PostForm("passwd" )
		//fmt.Print(passwd)
		isOkay := totp.Validate(passwd, "4INTFGQRRNLWKV4FMQSAUQTOBFSO6QU6")
		c.JSON(200, gin.H{
			"isOkay":isOkay,
		})
		//fmt.Print(isOkay)
	})
	router.Run(":8080")
}