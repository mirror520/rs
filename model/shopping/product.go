package shopping

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title string  // 名稱
	Price float64 // 定價
	Likes int     // 按讚數

	Styles []*ProductStyle // 樣式顏色
	Images []*ProductImage // 展示圖片
}

func (p *Product) UploadImage(url string) {
	image := &ProductImage{ImageURL: url}

	if p.Images == nil {
		p.Images = make([]*ProductImage, 0)
	}

	p.Images = append(p.Images, image)
}

func (p *Product) Like() {
	p.Likes++
}

type ProductStyle struct {
	ProductID uint   `gorm:"primarykey"`
	Style     string `gorm:"primarykey"`
}

type ProductImage struct {
	ProductID uint   `gorm:"primarykey"`
	ImageURL  string `gorm:"primarykey"`
}
