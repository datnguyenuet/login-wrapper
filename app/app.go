package app

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"login-wrapper/app/middleware"
)

var _ App = (*app)(nil)

type App interface {
	Routes(ctx context.Context) http.Handler
}

type app struct {
}

func New() App {
	return &app{}
}

func (a *app) Routes(ctx context.Context) http.Handler {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.SetRequestID())
	r.Use(middleware.SetLogger())

	r.GET("/health", a.handleHealth())
	r.POST("/auth/token", a.handleLogin())

	return r
}

func (a *app) handleError(ginCtx *gin.Context, err error) {
	ginCtx.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}
