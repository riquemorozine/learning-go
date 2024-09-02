package main

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    string `gorm:"primary_key"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?parseTime=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	products := []Product{
		{ID: uuid.New().String(), Name: "Laptop", Price: 2000},
		{ID: uuid.New().String(), Name: "Mouse", Price: 50},
		{ID: uuid.New().String(), Name: "Keyboard", Price: 100},
	}

	db.Create(&products)

	var p Product
	db.First(&p, "id = ?", products[0].ID)

	var pro []Product
	db.Limit(2).Offset(2).Find(&pro)

	for _, p := range pro {
		println(p.Name)
	}

	var newProducts []Product
	db.Where("price LIKE ?", "%book%").Find(&newProducts)
	for _, p := range newProducts {
		println(p.Name)
	}

	var pFirst Product
	db.First(&pFirst, 1)
	pFirst.Name = "New Mouse"
	db.Save(&pFirst)

	var p2 Product
	db.First(&p2, 2)
	fmt.Println(p2.Name)
	db.Delete(&p2)

}
