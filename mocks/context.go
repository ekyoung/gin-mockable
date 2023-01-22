package mocks

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type Context struct {
	mock.Mock
}

func (m *Context) Param(key string) string {
	ret := m.Called(key)
	return ret.String(0)
}

func (m *Context) HTML(code int, name string, obj interface{}) {
	m.Called(code, name, obj)
}

func (m *Context) String(code int, format string, values ...interface{}) {
	m.Called(code, format, values)
}

func (m *Context) Query(key string) string {
	ret := m.Called(key)
	return ret.String(0)
}

func (m *Context) AbortWithError(code int, err error) *gin.Error {
	ret := m.Called(code, err)
	return ret.Get(0).(*gin.Error)
}

func (m *Context) GetRequest() *http.Request {
	ret := m.Called()
	return ret.Get(0).(*http.Request)
}

func (m *Context) Redirect(code int, location string) {
	ret := m.Called(code, location)
}
