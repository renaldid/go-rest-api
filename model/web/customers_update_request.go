package web

type CustomersUpdateRequest struct {
	Id          int    `validate:"required"`
	Name        string `validate:"required,min=1,max=100" json:"name"`
	Address     string `validate:"required,min=1,max=250" json:"address"`
	Email       string `validate:"required,min=1,max=250" json:"email"`
	PhoneNumber string `validate:"required,min=1,max=100" json:"phoneNumber"`
}
