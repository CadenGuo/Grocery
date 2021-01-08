package db

import (
	"errors"
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
	result := db.DbIns.Select("GroceryItem").Delete(&GroceryItem{Id: id})
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
	query := db.DbIns.Where("")
	if id != nil {
		query = query.Where("id = ?", *id)
	}
	if name != nil {
		query = query.Where("name LIKE ?", "%"+*name+"%")
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
	result := query.Find(&groceryItems)
	if result.Error != nil {
		return []GroceryItem{}, result.Error
	}
	return groceryItems, nil
}

func (db *Manager) ListGroceryItemIdIn(ids []int) ([]GroceryItem, error) {
	var groceryItems []GroceryItem
	result := db.DbIns.Where("Id IN ?", ids).Find(&groceryItems)
	if result.Error != nil || len(groceryItems) != len(ids) {
		return []GroceryItem{}, errors.New("grocery item not found")
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
	query := db.DbIns.Where("")
	if id != nil {
		query = query.Where("id = ?", *id)
	}
	if name != nil {
		query = query.Where("name LIKE ?", "%"+*name+"%")
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
	result := query.Find(&drinks)
	if result.Error != nil {
		return []Drink{}, result.Error
	}
	return drinks, nil
}

func (db *Manager) CreateDishes(
	name string,
	complexity int,
	groceryItems []GroceryItem,
) (Dishes, error) {
	dishes := Dishes{Name: name, Complexity: complexity}
	err := db.DbIns.Create(&dishes).Association("GroceryItem").Append(&groceryItems)
	if err != nil {
		return Dishes{}, err
	}

	return dishes, nil
}

func (db *Manager) UpdateDishes(dishes Dishes, groceryItems *[]GroceryItem) (Dishes, error) {
	result := db.DbIns.Save(&dishes)
	if result.Error != nil {
		return Dishes{}, result.Error
	}

	if groceryItems != nil {
		err := db.DbIns.Model(&dishes).Association("GroceryItem").Delete(dishes.GroceryItem)
		if err != nil {
			return Dishes{}, err
		}
		err = result.Association("GroceryItem").Append(groceryItems)
		if err != nil {
			return Dishes{}, err
		}
	}

	return dishes, nil
}

func (db *Manager) DeleteDishes(id int) error {
	result := db.DbIns.Select("GroceryItem").Delete(&Dishes{Id: id})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *Manager) ListDishes(
	id *int,
	name *string,
	complexityLessThan *int,
	complexityMoreThan *int,
	groceryItemIds *[]int,
) ([]Dishes, error) {
	var dishesList []Dishes
	query := db.DbIns.Preload("GroceryItem")
	if id != nil {
		query = query.Where("id = ?", *id)
	}
	if name != nil {
		query = query.Where("name LIKE ?", "%"+*name+"%")
	}
	if complexityLessThan != nil {
		query = query.Where("complexity <= ?", complexityLessThan)
	}
	if complexityMoreThan != nil {
		query = query.Where("complexity >= ?", complexityMoreThan)
	}
	if groceryItemIds != nil {
		query = query.Joins("JOIN dishes_grocery_item ON dishes.id = dishes_grocery_item.dishes_id AND dishes_grocery_item.grocery_item_id in ?", *groceryItemIds)
	}
	result := query.Find(&dishesList)
	if result.Error != nil {
		return []Dishes{}, result.Error
	}
	return dishesList, nil
}
