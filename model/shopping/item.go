package shopping

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	Style     string
	StoreID   uint
	CouponID  uint

	Product      *Product      // 商品
	ProductStyle *ProductStyle `gorm:"foreignKey:ProductID,Style"` // 商品樣式顏色
	Store        *Store        // 賣場
	Coupon       *Coupon       // 使用折價券
}

func (i *Item) Price() float64 {
	var discount float64
	if i.Coupon != nil {
		discount = i.Coupon.Discount
	}
	return i.Product.Price - discount
}
