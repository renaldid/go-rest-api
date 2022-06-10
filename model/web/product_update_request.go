package web

type ProductsUpdateRequest struct {
	Id           int    `validate:"required"`
	Name         string `validate:"required,min=1,max=100" json:"name"`
	Price        int    `validate:"required,min=1,max=100" json:"price"`
	CategoryId   int    `validate:"required,min=1,max=100" json:"category_id"`
	CategoryName string `validate:"required,min=1,max=100" json:"category_name"`
}
