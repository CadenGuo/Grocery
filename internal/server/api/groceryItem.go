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
		var requestQuery ListGroceryItemSchema
		if err := ctx.ShouldBindQuery(&requestQuery); err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInvalidParameters))
			return
		}

		groceryItems, err := h.dbManager.ListGroceryItem(
			requestQuery.Id,
			requestQuery.Name,
			requestQuery.Type,
			requestQuery.ExpirationBefore,
			requestQuery.ExpirationAfter,
		)
		if err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInternalError))
			return
		}

		ctx.JSON(http.StatusOK, succeededResponse(groceryItems))

	case http.MethodPost:
		var requestBody CreateGroceryItemSchema
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInvalidParameters))
			return
		}

		_, err := h.dbManager.CreateGroceryItem(requestBody.Name, requestBody.Unit, requestBody.Amount, requestBody.Type, requestBody.Expiration)
		if err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInternalError))
			return
		}

		ctx.JSON(http.StatusOK, succeededResponse(common.EmptyMap))
	case http.MethodPut:
		var requestBody UpdateGroceryItemSchema
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInvalidParameters))
			return
		}

		groceryItems, err := h.dbManager.ListGroceryItem(
			&requestBody.Id,
			nil, nil, nil, nil,
		)
		if len(groceryItems) == 0 {
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInvalidParameters))
			return
		}
		if requestBody.Name != nil {
			groceryItems[0].Name = *requestBody.Name
		}
		if requestBody.Unit != nil {
			groceryItems[0].Unit = *requestBody.Unit
		}
		if requestBody.Amount != nil {
			groceryItems[0].Amount = *requestBody.Amount
		}
		if requestBody.Type != nil {
			groceryItems[0].Type = *requestBody.Type
		}
		if requestBody.Expiration != nil {
			groceryItems[0].Expiration = *requestBody.Expiration
		}

		_, err = h.dbManager.UpdateGroceryItem(groceryItems[0])
		if err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInternalError))
			return
		}
		ctx.JSON(http.StatusOK, succeededResponse(common.EmptyMap))
	case http.MethodDelete:
		var requestBody DeleteGroceryItemSchema
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInvalidParameters))
			return
		}

		err := h.dbManager.DeleteGroceryItem(requestBody.Id)
		if err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInternalError))
			return
		}

		ctx.JSON(http.StatusOK, succeededResponse(common.EmptyMap))
	}
}
