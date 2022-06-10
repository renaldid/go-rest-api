package web

type OrdersCreateRequest struct {
	CustomerId  int `validate:"required,min=1,max=11" json:"customer_id"`
	TotalAmount int `validate:"required,min=1,max=11" json:"total_amount"`
}
