package shopping

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerID uint
	StoreID    uint

	Customer *Customer // 顧客
	Store    *Store    // 賣場
	Items    []*Item   // 購買商品明細
}
