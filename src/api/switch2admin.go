package api

import (
	"github.com/gin-gonic/gin"
	"github.com/pquerna/otp/totp"
)

// Switch2Admin : api for switch to admin page
func Switch2Admin(c *gin.Context) {
	passwd := c.PostForm("passwd")
	//fmt.Print(passwd)
	isOkay := totp.Validate(passwd, "4INTFGQRRNLWKV4FMQSAUQTOBFSO6QU6")
	c.JSON(200, gin.H{
		"isOkay": isOkay,
	})
	//fmt.Print(isOkay)
}
