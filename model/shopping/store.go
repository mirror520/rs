package shopping

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Title string

	Coupons  []Coupon  // 折價券
	Products []Product // 商品型錄
}
