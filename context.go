package mockable

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Context interface {
	Param(key string) string
	HTML(code int, name string, obj interface{})
	String(code int, format string, values ...interface{})
	Query(key string) string
	AbortWithError(code int, err error) *gin.Error
	GetRequest() *http.Request
	SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
	Redirect(code int, location string)
}

type extendedGinContext struct {
	*gin.Context
}

func (c *extendedGinContext) GetRequest() *http.Request {
	return c.Request
}

type HandlerFunc func(Context)

func Handler(fn HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		fn(&extendedGinContext{
			c,
		})
	}
}
