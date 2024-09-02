package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	category := Category{Name: "Cozinha"}
	category2 := Category{Name: "Eletronicos"}

	db.Create(&category)
	db.Create(&category2)

	db.Create(&Product{
		Name:       "Geladeira",
		Price:      2000,
		Categories: []Category{category, category2},
	})

	db.Create(&SerialNumber{
		Number:    "123456",
		ProductId: 1,
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Product").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Product {
			fmt.Println("- ", product.Name)
		}
	}

	//category := Category{Name: "Eletronics"}
	//db.Create(&category)
	//
	//db.Create(&Product{
	//	Name:     "Laptop",
	//	Price:    2000,
	//	Category: category,
	//})
	//
	//db.Create(&SerialNumber{
	//	Number:    "123456",
	//	ProductId: 1,
	//})
	//
	//var products []Product
	//
	//db.Preload("Category").Preload("SerialNumber").Find(&products)
	//for _, product := range products {
	//	fmt.Println(product.Category.Name, product.SerialNumber.Number, product.Name)
	//}
	//
	//var categories []Category
	//
	//err = db.Model(&Category{}).Preload("Product.SerialNumber").Find(&categories).Error
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, category := range categories {
	//	fmt.Println(category.Name, ":")
	//	for _, product := range category.Product {
	//		fmt.Println(product.Name, product.SerialNumber.Number)
	//	}
	//}
}
