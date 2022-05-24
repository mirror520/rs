package model

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Title string

	Coupn    []Coupn   // 折價券
	Products []Product // 商品型錄
}
