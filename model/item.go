package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ProductID uint
	StoreID   uint
	CoupnID   uint

	Product Product // 商品
	Store   Store   // 賣場
	Coupn   Coupn   // 使用折價券
}
