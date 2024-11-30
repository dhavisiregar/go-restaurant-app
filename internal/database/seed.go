package database

import (
	"github.com/dhavisiregar/go-restaurant-app/internal/model"
	"github.com/dhavisiregar/go-restaurant-app/internal/model/constant"
	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	db.AutoMigrate(&model.MenuItem{}, &model.Order{}, &model.ProductOrder{}, &model.User{})

	foodMenu := []model.MenuItem{
		{
			Name:      "Chicken Wings",
            OrderCode: "chicken_wings",
            Price:      37500,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Spaghetti Carbonara",
            OrderCode: "spaghetti_carbonara",
            Price:      50000,
			Type:      constant.MenuTypeFood,
		},
	}

	drinkMenu := []model.MenuItem{
		{
            Name:      "Orange Juice",
            OrderCode: "orange_juice",
            Price:      10000,
			Type:      constant.MenuTypeDrink,
        },
        {
            Name:      "Ice Tea",
            OrderCode: "ice_tea",
            Price:      5000,
			Type:      constant.MenuTypeDrink,
        },
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}
	
}