package main

import (
	"fmt"
	"testing"

	"github.com/mirror520/rs/model"
)

func TestModel(t *testing.T) {
	store := model.Store{
		Title: "犀牛盾",
		Coupns: []model.Coupn{
			{Title: "現折 $100", Discount: -100},
			{Title: "現折 $200", Discount: -200},
			{Title: "現折 $300", Discount: -300},
		},
	}

	product := model.Product{
		Title: "犀牛盾 適用Airpods Pro/第2代/第1代 防摔保護套(含扣環)/玩具總動員-倒吊三眼怪",
		Price: 690,
		Likes: 609,
		Styles: []model.ProductStyle{
			{Style: "白"},
			{Style: "雀藍"},
			{Style: "石墨黑"},
			{Style: "薄荷綠"},
			{Style: "鵝黃"},
			{Style: "櫻花粉"},
		},
		Images: []model.ProductImage{
			{ImageURL: "../image1.jpg"},
			{ImageURL: "../image2.jpg"},
			{ImageURL: "../image3.jpg"},
			{ImageURL: "../image4.jpg"},
			{ImageURL: "../image5.jpg"},
		},
	}

	customer := model.Customer{
		Name: "Mirror",
	}

	order := model.Order{
		Customer: customer,
		Store:    store,
		Items: []model.Item{
			{
				Product: product,
				Store:   store,
				Coupn:   store.Coupns[0],
			},
		},
	}

	fmt.Println(order)
}
