package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"grocery/internal/common"
	"net/http"
)

func (h *Handler) GroceryItem(ctx *gin.Context) {
	h.LogRequest(ctx)
	method := ctx.Request.Method

	switch method {
	case http.MethodGet:
		log.Error().Msg("error")
		ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInternalError))
	case http.MethodPost:
		ctx.JSON(http.StatusOK, succeededResponse(gin.H{"status": "create item"}))
	case http.MethodPut:
		ctx.JSON(http.StatusOK, succeededResponse(gin.H{"status": "update item"}))
	}
}

func (h *Handler) Dishes(ctx *gin.Context) {
	h.LogRequest(ctx)
	method := ctx.Request.Method

	switch method {
	case http.MethodPost:
		log.Error().Msg("error")
		ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInternalError))
	case http.MethodGet:
		ctx.JSON(http.StatusOK, succeededResponse(gin.H{"status": "get dishes"}))
	}
}
