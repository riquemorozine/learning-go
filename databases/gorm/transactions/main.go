package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category struct {
	ID      int `gorm:"primaryKey"`
	Name    string
	Product []Product `gorm:"many2many:product_categories;"`
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	Categories   []Category `gorm:"many2many:product_categories;"`
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductId int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Category{}, &Product{}, &SerialNumber{})

	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	c.Name = "Sla"
	tx.Debug().Save(&c)
	tx.Commit()
}
