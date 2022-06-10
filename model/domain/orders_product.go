package domain

import "time"

type OrdersProduct struct {
	Id         int
	OrderId    int
	ProductId  int
	Qty        int
	Amount     int
	CreatedAt  time.Time
	UploadedAt time.Time
}
