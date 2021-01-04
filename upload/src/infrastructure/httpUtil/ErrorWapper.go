package httpUtil

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type HandlerFunc func(c *gin.Context) error

func ErrorWrapper(handler HandlerFunc, errorMapper func(error) string) func(c *gin.Context) {
	return func(c *gin.Context) {
		err := handler(c)
		if err != nil {
			log.Println(err)
			_ = c.AbortWithError(http.StatusBadRequest, errors.New(errorMapper(err)))
		}
	}
}
