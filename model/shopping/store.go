package shopping

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Title string

	Coupons  []*Coupon  // 折價券
	Products []*Product `gorm:"many2many:items"` // 商品型錄
}

func (s *Store) IssueCoupon(title string, discount float64) {
	coupon := &Coupon{
		Title:    title,
		Discount: discount,
	}

	if s.Coupons == nil {
		s.Coupons = make([]*Coupon, 0)
	}

	s.Coupons = append(s.Coupons, coupon)
}
