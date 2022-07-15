package domain

type OrdersProduct struct {
	Id         int
	OrderId    int
	ProductId  int
	Qty        int
	Amount     int
	CreatedAt  string
	UploadedAt string
}
