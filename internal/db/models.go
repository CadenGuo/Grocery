package db

import (
	"database/sql/driver"
	"time"
)

type GroceryItemType string

const (
	Protein     GroceryItemType = "protein"
	Vegetable   GroceryItemType = "vegetable"
	Fruit       GroceryItemType = "fruit"
	Cereals     GroceryItemType = "cereals"
	InstantFood GroceryItemType = "instant_food"
)

func (e *GroceryItemType) Scan(value interface{}) error {
	*e = GroceryItemType(value.([]byte))
	return nil
}

func (e GroceryItemType) Value() (driver.Value, error) {
	return string(e), nil
}

type GroceryItem struct {
	Id         int             `json:"id" gorm:"type:int;autoIncrement;not null"`
	Name       string          `json:"name" gorm:"size:100;not null"`
	Unit       string          `json:"unit" gorm:"type:text;size:100"`
	Amount     float64         `json:"amount" gorm:"type:float"`
	Type       GroceryItemType `json:"type" gorm:"type:ENUM('protein', 'vegetable', 'fruit', 'cereals', 'instant_food')"`
	Expiration time.Time       `json:"expiration"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
}

func (GroceryItem) TableName() string {
	return "grocery_item"
}

type Drink struct {
	Id         int       `json:"id" gorm:"type:int;autoIncrement;not null"`
	Name       string    `json:"name" gorm:"size:100;not null"`
	Unit       string    `json:"unit" gorm:"type:text;size:100"`
	Amount     float64   `json:"amount" gorm:"type:float"`
	Expiration time.Time `json:"expiration"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Carbonated bool      `json:"carbonated"`
	Alcoholic  bool      `json:"alcoholic"`
}

func (Drink) TableName() string {
	return "drink"
}

type Dishes struct {
	Id          int           `json:"id" gorm:"type:int;autoIncrement;not null"`
	Name        string        `json:"name" gorm:"size:100;not null"`
	GroceryItem []GroceryItem `json:"grocery_item" gorm:"many2many:dishes_grocery_item;"`
	Complexity  int           `json:"complexity" gorm:"check:complexity > 0"`
}

func (Dishes) TableName() string {
	return "dishes"
}
