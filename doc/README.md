# Data Modeling

## Requirements

- 一個賣場 `Store` 可以上架多件產品 `Products`，可販售產品視為現貨 `Goods`。
- 一件產品 `Product` 可以在多個賣場 `Stores` 販售，賣場包含實體商店與網路平台。
- 每件產品 `Product` 包含產品名稱 `title`、定價 `price`、樣式顏色 `styles`、展示圖片 `images`、及顧客點讚數 `Likes`。
- 每個賣場 `Store` 會發行 (`issue`) 自己的折價券 `Coupon`。
- 顧客 `Consumer` 可以使用 (`use`) 賣場折價券，對購買產品定價進行折扣。
- 顧客 `Consumer` 可以對賣場 `Store` 下訂單 `Order`。
- 訂單 `Order` 會紀錄所購買的產品現貨 `Goods` 明細，且每件現貨可以單獨使用賣場折價券 `Coupon`。
