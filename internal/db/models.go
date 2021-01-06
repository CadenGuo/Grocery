package db

import "time"

type Protein struct {
	Id         int     `gorm:"type:int;autoIncrement;not null"`
	Name       string  `gorm:"size:100;not null"`
	Unit       string  `gorm:"type:text;size:100"`
	Amount     float64 `gorm:"type:float"`
	Expiration time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (Protein) TableName() string {
	return "protein"
}

type Vegetable struct {
	Id         int     `gorm:"type:int;autoIncrement;not null"`
	Name       string  `gorm:"size:100;not null"`
	Unit       string  `gorm:"type:text;size:100"`
	Amount     float64 `gorm:"type:float"`
	Expiration time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (Vegetable) TableName() string {
	return "vegetable"
}

type Fruit struct {
	Id         int     `gorm:"type:int;autoIncrement;not null"`
	Name       string  `gorm:"size:100;not null"`
	Unit       string  `gorm:"type:text;size:100"`
	Amount     float64 `gorm:"type:float"`
	Expiration time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (Fruit) TableName() string {
	return "fruit"
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

type Cereals struct {
	Id         int     `gorm:"type:int;autoIncrement;not null"`
	Name       string  `gorm:"size:100;not null"`
	Unit       string  `gorm:"type:text;size:100"`
	Amount     float64 `gorm:"type:float"`
	Expiration time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (Cereals) TableName() string {
	return "cereals"
}

type InstantFood struct {
	Id         int     `gorm:"type:int;autoIncrement;not null"`
	Name       string  `gorm:"size:100;not null"`
	Unit       string  `gorm:"type:text;size:100"`
	Amount     float64 `gorm:"type:float"`
	Expiration time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (InstantFood) TableName() string {
	return "instant_food"
}

type Dishes struct {
	Id         int         `gorm:"type:int;autoIncrement;not null"`
	Name       string      `gorm:"size:100;not null"`
	Protein    []Protein   `gorm:"many2many:dishes_protein;"`
	Vegetable  []Vegetable `gorm:"many2many:dishes_vegetable;"`
	Fruit      []Fruit     `gorm:"many2many:dishes_fruit;"`
	Cereals    []Cereals   `gorm:"many2many:dishes_cereals;"`
	Complexity int         `gorm:"check:complexity > 0"`
}

func (Dishes) TableName() string {
	return "dishes"
}
