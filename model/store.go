package model

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Title string

	Coupns   []Coupn   // 折價券
	Products []Product // 商品型錄
}
