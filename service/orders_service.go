package service

import (
	"context"
	"go_rest_api/model/web"
)

type OrdersService interface {
	Create(ctx context.Context, request web.OrdersCreateRequest) web.OrdersResponse
	Update(ctx context.Context, request web.OrdersUpdateRequest) web.OrdersResponse
	Delete(ctx context.Context, ordersId int)
	FindById(ctx context.Context, ordersId int) web.OrdersResponse
	FindByAll(ctx context.Context) []web.OrdersResponse
}
