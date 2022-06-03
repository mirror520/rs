package shopping

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	StoreID   uint
	CouponID  uint

	Product *Product // 商品
	Store   *Store   // 賣場
	Coupon  *Coupon  // 使用折價券
}
