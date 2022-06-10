package service

import (
	"context"
	"go_rest_api/model/web"
)

type CustomersService interface {
	Create(ctx context.Context, request web.CustomersCreateRequest) web.CustomersResponse
	Update(ctx context.Context, request web.CustomersUpdateRequest) web.CustomersResponse
	Delete(ctx context.Context, customersId int)
	FindById(ctx context.Context, customersId int) web.CustomersResponse
	FindByAll(ctx context.Context) []web.CustomersResponse
}
