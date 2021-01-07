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
	Id         int             `gorm:"type:int;autoIncrement;not null"`
	Name       string          `gorm:"size:100;not null"`
	Unit       string          `gorm:"type:text;size:100"`
	Amount     float64         `gorm:"type:float"`
	Expiration time.Time
	Type       GroceryItemType `json:"type" gorm:"type:ENUM('protein', 'vegetable', 'fruit', 'cereals', 'instant_food')"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (GroceryItem) TableName() string {
	return "grocery_item"
}

type Drink struct {
	Id         int     `gorm:"type:int;autoIncrement;not null"`
	Name       string  `gorm:"size:100;not null"`
	Unit       string  `gorm:"type:text;size:100"`
	Amount     float64 `gorm:"type:float"`
	Expiration time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Carbonated bool
	Alcoholic  bool
}

func (Drink) TableName() string {
	return "drink"
}

type Dishes struct {
	Id         int           `gorm:"type:int;autoIncrement;not null"`
	Name       string        `gorm:"size:100;not null"`
	Protein    []GroceryItem `gorm:"many2many:dishes_grocery_item;"`
	Complexity int           `gorm:"check:complexity > 0"`
}

func (Dishes) TableName() string {
	return "dishes"
}
