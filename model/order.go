package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerID string
	StoreID    string

	Customer Customer // 顧客
	Store    Store    // 賣場
	Items    []Item   // 購買商品明細
}
