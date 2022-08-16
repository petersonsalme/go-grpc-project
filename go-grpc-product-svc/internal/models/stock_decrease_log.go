package models

type StockDecreaseLog struct {
	ID        int64 `json:"id" gorm:"primaryKey"`
	OrderID   int64 `json:"order_id"`
	ProductID int64 `json:"product_id"`
}
