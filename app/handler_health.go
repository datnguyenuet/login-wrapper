package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"login-wrapper/pkg/logging"
)

func (a *app) handleHealth() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		ctx := ginCtx.Request.Context()
		log := logging.FromContext(ctx)
		log.Infof("starting handle health api ...")

		ginCtx.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
		return
	}
}
