package shopping

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name string // 顧客名稱

	Orders []Order // 顧客所有訂單
}
