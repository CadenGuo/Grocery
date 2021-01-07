package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"grocery/internal/common"
	"net/http"
)

func (h *Handler) Drink(ctx *gin.Context) {
	h.LogRequest(ctx)
	method := ctx.Request.Method

	switch method {
	case http.MethodGet:
		var requestQuery ListDrinkSchema
		if err := ctx.ShouldBindQuery(&requestQuery); err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInvalidParameters))
			return
		}

		drinks, err := h.dbManager.ListDrink(
			requestQuery.Id,
			requestQuery.Name,
			requestQuery.Carbonated,
			requestQuery.Alcoholic,
			requestQuery.ExpirationBefore,
			requestQuery.ExpirationAfter,
		)
		if err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInternalError))
			return
		}

		ctx.JSON(http.StatusOK, succeededResponse(drinks))

	case http.MethodPost:
		var requestBody CreateDrinkSchema
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInvalidParameters))
			return
		}

		_, err := h.dbManager.CreateDrink(
			requestBody.Name,
			requestBody.Unit,
			requestBody.Amount,
			requestBody.Carbonated,
			requestBody.Alcoholic,
			requestBody.Expiration,
		)
		if err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInternalError))
			return
		}

		ctx.JSON(http.StatusOK, succeededResponse(common.EmptyMap))
	case http.MethodPut:
		var requestBody UpdateDrinkSchema
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInvalidParameters))
			return
		}

		drinks, err := h.dbManager.ListDrink(
			&requestBody.Id,
			nil, nil, nil, nil, nil,
		)
		if len(drinks) == 0 {
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInvalidParameters))
			return
		}
		if requestBody.Name != nil {
			drinks[0].Name = *requestBody.Name
		}
		if requestBody.Unit != nil {
			drinks[0].Unit = *requestBody.Unit
		}
		if requestBody.Amount != nil {
			drinks[0].Amount = *requestBody.Amount
		}
		if requestBody.Carbonated != nil {
			drinks[0].Carbonated = *requestBody.Carbonated
		}
		if requestBody.Alcoholic != nil {
			drinks[0].Alcoholic = *requestBody.Alcoholic
		}
		if requestBody.Expiration != nil {
			drinks[0].Expiration = *requestBody.Expiration
		}

		_, err = h.dbManager.UpdateDrink(drinks[0])
		if err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInternalError))
			return
		}
		ctx.JSON(http.StatusOK, succeededResponse(common.EmptyMap))
	case http.MethodDelete:
		var requestBody DeleteDrinkSchema
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInvalidParameters))
			return
		}

		err := h.dbManager.DeleteDrink(requestBody.Id)
		if err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInternalError))
			return
		}

		ctx.JSON(http.StatusOK, succeededResponse(common.EmptyMap))
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
