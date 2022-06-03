package persistent

import (
	"github.com/mirror520/rs/model/shopping"
)

func Seed(repo shopping.Repository) error {
	// 建立新的賣場
	store := &shopping.Store{
		Title: "犀牛盾",
	}
	if err := repo.CreateStore(store); err != nil {
		return err
	}

	// 賣場發行折價券
	store.IssueCoupon("現折 $100", 100)
	store.IssueCoupon("現折 $200", 200)
	store.IssueCoupon("現折 $300", 300)
	if err := repo.UpdateStore(store); err != nil {
		return err
	}

	// 建立新商品資料
	product := &shopping.Product{
		Title: "犀牛盾 適用Airpods Pro/第2代/第1代 防摔保護套(含扣環)/玩具總動員-倒吊三眼怪",
		Price: 690,
		Likes: 0,
		Styles: []*shopping.ProductStyle{
			{Style: "白"},
			{Style: "雀藍"},
			{Style: "石墨黑"},
			{Style: "薄荷綠"},
			{Style: "鵝黃"},
			{Style: "櫻花粉"},
		},
	}
	if err := repo.CreateProduct(product); err != nil {
		return err
	}

	// 上傳多張商品展示圖片
	product.UploadImage("../image1.jpg")
	product.UploadImage("../image2.jpg")
	product.UploadImage("../image3.jpg")
	product.UploadImage("../image4.jpg")
	product.UploadImage("../image5.jpg")

	// 顧客按讚 609 次
	for i := 0; i < 609; i++ {
		product.Like()
	}
	if err := repo.UpdateProduct(product); err != nil {
		return err
	}

	// 新的顧客註冊
	customer := &shopping.Customer{
		Name: "Mirror",
	}
	if err := repo.RegisterCustomer(customer); err != nil {
		return err
	}

	// 下訂單
	order := &shopping.Order{
		Customer: customer,
		Store:    store,
		Items: []*shopping.Item{
			{
				Product:      product,
				ProductStyle: product.Styles[0],
				Store:        store,
				Coupon:       store.Coupons[0],
			},
		},
	}
	if err := repo.PlaceOrder(order); err != nil {
		return err
	}

	return nil
}
