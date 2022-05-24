package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title string  // 名稱
	Price float64 // 定價
	Likes int64   // 按讚數

	Styles []ProductStyle // 樣式顏色
	Images []ProductImage // 展示圖片
	Stores []Store        // 鋪貨賣場
}

type ProductStyle struct {
	ProductID uint   `gorm:"primarykey"`
	Style     string `gorm:"primarykey"`
}

type ProductImage struct {
	ProductID uint   `gorm:"primarykey"`
	ImageURL  string `gorm:"primarykey"`
}
