package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// db, err := gorm.Open("sqlite3", "test.db")
	// db, err := gorm.Open("mysql", "root:1234567890@tcp(127.0.0.1:3306)/xcar?charset=utf8&parseTime=true")
	db, err := gorm.Open("mysql", "root:1234567890@(127.0.0.1:3306)/xcardb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
		// root:1234567890@tcp(127.0.0.1:3306)/xcar?charset=utf8&parseTime=true
	}
	defer db.Close()

	// 自动迁移模式
	db.AutoMigrate(&Product{})

	// 创建
	db.Create(&Product{Code: "L1212", Price: 1000})

	// 读取
	var product Product
	db.First(&product, 1)                   // 查询id为1的product
	db.First(&product, "code = ?", "L1212") // 查询code为l1212的product
	fmt.Println(product)
	// 更新 - 更新product的price为2000
	db.Model(&product).Update("Price", 2000)
	fmt.Println(product)
	// 删除 - 删除product
	db.Delete(&product)
}
