package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestHttpServer(t *testing.T) {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "test",
		})
	})
	r.GET("/test/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "api",
		})
	})
	r.Run("127.0.0.1:8080")
}
