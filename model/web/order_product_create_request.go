package web

type OrderProductCreateRequest struct {
	OrderId    int    `validate:"required" json:"order_id"`
	ProductId  int    `validate:"required" json:"product_id"`
	Qty        int    `validate:"required" json:"qty"`
	Amount     int    `validate:"required" json:"amount"`
	CreatedAt  string `validate:"required,min=1max=100" json:"created_at"`
	UploadedAt string `validate:"required,min=1max=100" json:"uploaded_at"`
}
