package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func M() gin.HandlerFunc {
	return func(c *gin.Context) {
		H := c.Request.Host
		fmt.Println("ddd")
		c.Set("M", gin.H{"H":H, "Ln":Ln(H)})
		c.Next()
	}
}

func Ln(domain string) string {
	return "cn"

	// Todo: Select Languages
	//sp := strings.Split(domain, ":")
	//switch sp[0] {
	//case "cn":
	//	return "zh"
	//default:
	//	return "zh"
	//}
}