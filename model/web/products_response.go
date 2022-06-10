package web

type ProductsResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}
