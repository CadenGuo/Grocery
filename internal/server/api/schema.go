package api

import (
	"grocery/internal/db"
	"time"
)

type CreateGroceryItemSchema struct {
	Name       string             `json:"name" binding:"required"`
	Unit       string             `json:"unit" binding:"required"`
	Amount     float64            `json:"amount" binding:"required"`
	Type       db.GroceryItemType `json:"type" binding:"required"`
	Expiration *time.Time         `json:"expiration,omitempty"`
}

type UpdateGroceryItemSchema struct {
	Id         int                 `json:"id" binding:"required"`
	Name       *string             `json:"name"`
	Unit       *string             `json:"unit"`
	Amount     *float64            `json:"amount"`
	Type       *db.GroceryItemType `json:"type"`
	Expiration *time.Time          `json:"expiration,omitempty"`
}

type DeleteGroceryItemSchema struct {
	Id int `json:"id" binding:"required"`
}

type ListGroceryItemSchema struct {
	Id               *int                `form:"id"`
	Name             *string             `form:"name"`
	Type             *db.GroceryItemType `form:"type"`
	ExpirationBefore *time.Time          `form:"expiration_before"`
	ExpirationAfter  *time.Time          `form:"expiration_after"`
}

type CreateDrinkSchema struct {
	Name       string     `json:"name" binding:"required"`
	Unit       string     `json:"unit" binding:"required"`
	Amount     float64    `json:"amount" binding:"required"`
	Carbonated bool       `json:"carbonated"`
	Alcoholic  bool       `json:"alcoholic"`
	Expiration *time.Time `json:"expiration,omitempty"`
}

type UpdateDrinkSchema struct {
	Id         int        `json:"id" binding:"required"`
	Name       *string    `json:"name"`
	Unit       *string    `json:"unit"`
	Amount     *float64   `json:"amount"`
	Carbonated *bool      `json:"carbonated"`
	Alcoholic  *bool      `json:"alcoholic"`
	Expiration *time.Time `json:"expiration,omitempty"`
}

type ListDrinkSchema struct {
	Id               *int       `form:"id"`
	Name             *string    `form:"name"`
	Carbonated       *bool      `form:"carbonated"`
	Alcoholic        *bool      `form:"alcoholic"`
	ExpirationBefore *time.Time `form:"expiration_before"`
	ExpirationAfter  *time.Time `form:"expiration_after"`
}

type DeleteDrinkSchema struct {
	Id int `json:"id" binding:"required"`
}

type CreateDishesSchema struct {
	Name           string `json:"name" binding:"required"`
	Complexity     int    `json:"complexity"`
	GroceryItemIds []int  `json:"grocery_item_ids" binding:"required"`
}

type UpdateDishesSchema struct {
	Id             int     `json:"id" binding:"required"`
	Name           *string `json:"name"`
	Complexity     *int    `json:"complexity"`
	GroceryItemIds *[]int  `json:"grocery_item_ids"`
}

type ListDishesSchema struct {
	Id                 *int    `form:"id"`
	Name               *string `form:"name"`
	ComplexityLessThan *int    `form:"complexity_less_than"`
	ComplexityMoreThan *int    `form:"complexity_more_than"`
	GroceryItemIds     *[]int  `form:"grocery_item_ids"`
}

type DeleteDishesSchema struct {
	Id int `json:"id" binding:"required"`
}
