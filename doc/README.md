# Data Model

## Requirements (需求分析)

- 一個賣場 `Store` 可以上架多件產品 `Products`，可販售產品視為現貨 `Goods`。
- 一件產品 `Product` 可以在多個賣場 `Stores` 販售，賣場包含實體商店與網路平台。
- 每件產品 `Product` 包含產品名稱 `title`、定價 `price`、樣式顏色 `styles`、展示圖片 `images`、及顧客點讚數 `Likes`。
- 每個賣場 `Store` 會發行 (`issue`) 自己的折價券 `Coupon`。
- 顧客 `Consumer` 可以使用 (`use`) 賣場折價券，對購買產品定價進行折扣。
- 顧客 `Consumer` 可以對賣場 `Store` 下訂單 `Order`。
- 訂單 `Order` 會紀錄所購買的產品現貨 `Goods` 明細，且每件現貨可以單獨使用賣場折價券 `Coupon`。

## Data Modeling (資料塑模)

此為概念設計，產出 `ER Model`

![ER Model](./model.png)

註：
- 資料塑模是將真實世界按著需求抽象化塑模而成，所以模型中不需包含實體欄位細節。
- 裡面有使用到實體關係圖 (長方形 + 菱形)，表達其既是實體 `Entity` 又是關係 `Relationship`，其可表現實體關係建立的先後順序，例如無賣場或無顧客，則並無法下訂單，且訂單並不會憑空產生，故其表達力較強，不需依靠弱實體。

## 邏輯設計與實體設計

- 邏輯設計詳 `ORM` (../model)

- 實體設計由 `ORM` 處理，不呈現。
