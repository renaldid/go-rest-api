package web

type OrdersResponse struct {
	Id          int `json:"id"`
	CustomerId  int `json:"customer_id"`
	TotalAmount int `json:"total_amount"`
}
