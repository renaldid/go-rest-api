package web

type OrderProductResponse struct {
	Id         int    `json:"id"`
	OrderId    int    `json:"order_id"`
	ProductId  int    `json:"product_id"`
	Qty        int    `json:"qty"`
	Amount     int    `json:"amount"`
	CreatedAt  string `json:"created_at"`
	UploadedAt string `json:"uploaded_at"`
}
