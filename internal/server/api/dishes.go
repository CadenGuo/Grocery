package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"grocery/internal/common"
	"net/http"
)

func (h *Handler) Dishes(ctx *gin.Context) {
	h.LogRequest(ctx)
	method := ctx.Request.Method

	switch method {
	case http.MethodGet:
		var requestQuery ListDishesSchema
		if err := ctx.ShouldBindQuery(&requestQuery); err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInvalidParameters))
			return
		}

		fmt.Println(requestQuery)

		dishes, err := h.dbManager.ListDishes(
			requestQuery.Id,
			requestQuery.Name,
			requestQuery.ComplexityLessThan,
			requestQuery.ComplexityMoreThan,
			requestQuery.GroceryItemIds,
		)
		if err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInternalError))
			return
		}

		ctx.JSON(http.StatusOK, succeededResponse(dishes))

	case http.MethodPost:
		var requestBody CreateDishesSchema
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInvalidParameters))
			return
		}

		groceryItems, err := h.dbManager.ListGroceryItemIdIn(requestBody.GroceryItemIds)
		if err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInvalidParameters))
			return
		}

		_, err = h.dbManager.CreateDishes(
			requestBody.Name,
			requestBody.Complexity,
			groceryItems,
		)
		if err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInternalError))
			return
		}

		ctx.JSON(http.StatusOK, succeededResponse(common.EmptyMap))
	case http.MethodPut:
		var requestBody UpdateDishesSchema
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInvalidParameters))
			return
		}

		dishesList, err := h.dbManager.ListDishes(
			&requestBody.Id,
			nil, nil, nil, nil,
		)
		if len(dishesList) == 0 {
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInvalidParameters))
			return
		}
		if requestBody.Name != nil {
			dishesList[0].Name = *requestBody.Name
		}
		if requestBody.Complexity != nil {
			dishesList[0].Complexity = *requestBody.Complexity
		}
		if requestBody.GroceryItemIds != nil {
			groceryItems, err := h.dbManager.ListGroceryItemIdIn(*requestBody.GroceryItemIds)
			if err != nil {
				log.Warn().Msg(err.Error())
				ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInvalidParameters))
				return
			}
			_, err = h.dbManager.UpdateDishes(dishesList[0], &groceryItems)
		} else {
			_, err = h.dbManager.UpdateDishes(dishesList[0], nil)
		}

		if err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInternalError))
			return
		}
		ctx.JSON(http.StatusOK, succeededResponse(common.EmptyMap))
	case http.MethodDelete:
		var requestBody DeleteDishesSchema
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInvalidParameters))
			return
		}

		err := h.dbManager.DeleteDishes(requestBody.Id)
		if err != nil {
			log.Warn().Msg(err.Error())
			ctx.JSON(http.StatusOK, failedResponse(common.EmptyMap, common.ErrorInternalError))
			return
		}

		ctx.JSON(http.StatusOK, succeededResponse(common.EmptyMap))
	}
}
