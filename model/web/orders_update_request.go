package web

type OrdersUpdateRequest struct {
	Id          int `validate:"required"`
	CustomerId  int `validate:"required,min=1,max=11" json:"customer_id"`
	TotalAmount int `validate:"required,min=1,max=11" json:"total_amount"`
}
