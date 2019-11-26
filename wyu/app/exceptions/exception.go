package exceptions

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Exceptions struct {

}

func NewExceptions() *Exceptions {
	return &Exceptions{}
}

func (exp *Exceptions) NoRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"status":404})
}

func (exp *Exceptions) NoMethod(c *gin.Context) {

}
