package middleware

import (
	"github.com/gin-gonic/gin"
)

/**
 * Todo: Middleware Initialized
**/
func M() gin.HandlerFunc {
	return func(c *gin.Context) {
		H := c.Request.Host

		c.Set("M", gin.H{"H":H, "Ln":Ln(H)})
		c.Next()
	}
}

/**
 * Todo: Select Languages
**/
func Ln(domain string) (str string) {
	str = "cn"
	return
}