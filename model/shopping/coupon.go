package shopping

import "gorm.io/gorm"

type Coupon struct {
	gorm.Model
	Title    string  // 折價券
	Discount float64 // 折價金額
	StoreID  uint
}
