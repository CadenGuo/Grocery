package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"grocery/config"
	"grocery/internal/common"
	"time"
)

type Manager struct {
	DbIns *gorm.DB
}

func (db *Manager) Connect(c config.Conf) error {
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v)/%v?charset=utf8&parseTime=true",
		c.Db.User,
		c.Db.Password,
		c.Db.Host,
		c.Db.Name,
	)
	dbInstance, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	dbInstance.Debug()

	if err != nil {
		return err
	}
	db.DbIns = dbInstance
	return nil
}

func (db *Manager) Migrate() error {
	return db.DbIns.AutoMigrate(
		&GroceryItem{},
		&Dishes{},
		&Drink{},
	)
}

func getExpiration(expiration *time.Time) time.Time {
	var expirationValue time.Time
	if expiration != nil {
		expirationValue = *expiration
	} else {
		expirationValue = common.DefaultExpiration
		fmt.Println(expiration)
	}
	return expirationValue
}

func (db *Manager) CreateGroceryItem(
	name string,
	unit string,
	amount float64,
	itemType GroceryItemType,
	expiration *time.Time,
) (GroceryItem, error) {
	groceryItem := GroceryItem{Name: name, Unit: unit, Amount: amount, Type: itemType, Expiration: getExpiration(expiration)}
	result := db.DbIns.Create(&groceryItem)
	if result.Error != nil {
		return GroceryItem{}, result.Error
	}
	return groceryItem, nil
}

func (db *Manager) UpdateGroceryItem(groceryItem GroceryItem) (GroceryItem, error) {
	result := db.DbIns.Save(&groceryItem)
	if result.Error != nil {
		return GroceryItem{}, result.Error
	}
	return groceryItem, nil
}

func (db *Manager) DeleteGroceryItem(id int) error {
	result := db.DbIns.Delete(&GroceryItem{Id: id})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *Manager) ListGroceryItem(
	id *int,
	name *string,
	itemType *GroceryItemType,
	expirationBefore *time.Time,
	expirationAfter *time.Time,
) ([]GroceryItem, error) {
	var groceryItems []GroceryItem
	var result *gorm.DB
	query := db.DbIns.Where("")
	if id != nil {
		query = query.Where("id = ?", *id)
	}
	if name != nil {
		query = query.Where("name = ?", *name)
	}
	if itemType != nil {
		query = query.Where("type = ?", *itemType)
	}
	if expirationBefore != nil {
		query = query.Where("expiration < ?", *expirationBefore)
	}
	if expirationAfter != nil {
		query = query.Where("expiration > ?", *expirationAfter)
	}
	result = query.Find(&groceryItems)
	if result.Error != nil {
		return []GroceryItem{}, result.Error
	}
	return groceryItems, nil
}

func (db *Manager) CreateDrink(
	name string,
	unit string,
	amount float64,
	carbonated bool,
	alcoholic bool,
	expiration *time.Time,
) (Drink, error) {
	drink := Drink{Name: name, Unit: unit, Amount: amount, Carbonated: carbonated, Alcoholic: alcoholic, Expiration: getExpiration(expiration)}
	result := db.DbIns.Create(&drink)
	if result.Error != nil {
		return Drink{}, result.Error
	}
	return drink, nil
}

func (db *Manager) UpdateDrink(drink Drink) (Drink, error) {
	result := db.DbIns.Save(&drink)
	if result.Error != nil {
		return Drink{}, result.Error
	}
	return drink, nil
}

func (db *Manager) DeleteDrink(id int) error {
	result := db.DbIns.Delete(&Drink{Id: id})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *Manager) ListDrink(
	id *int,
	name *string,
	carbonated *bool,
	alcoholic *bool,
	expirationBefore *time.Time,
	expirationAfter *time.Time,
) ([]Drink, error) {
	var drinks []Drink
	var result *gorm.DB
	query := db.DbIns.Where("")
	if id != nil {
		query = query.Where("id = ?", *id)
	}
	if name != nil {
		query = query.Where("name = ?", *name)
	}
	if carbonated != nil {
		query = query.Where("carbonated = ?", *carbonated)
	}
	if alcoholic != nil {
		query = query.Where("alcoholic = ?", *alcoholic)
	}
	if expirationBefore != nil {
		query = query.Where("expiration < ?", *expirationBefore)
	}
	if expirationAfter != nil {
		query = query.Where("expiration > ?", *expirationAfter)
	}
	result = query.Find(&drinks)
	if result.Error != nil {
		return []Drink{}, result.Error
	}
	return drinks, nil
}
