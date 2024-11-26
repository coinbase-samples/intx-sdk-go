/**
 * Copyright 2024-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package orders

import (
	"context"

	"github.com/coinbase-samples/intx-sdk-go/client"
)

type OrdersService interface {
	CancelOrder(ctx context.Context, request *CancelOrderRequest) (*CancelOrderResponse, error)
	CancelOrders(ctx context.Context, request *CancelOrdersRequest) (*CancelOrdersResponse, error)
	CreateOrder(ctx context.Context, request *CreateOrderRequest) (*CreateOrderResponse, error)
	GetOrderDetails(ctx context.Context, request *GetOrderDetailsRequest) (*GetOrderDetailsResponse, error)
	ListOpenOrders(ctx context.Context, request *ListOpenOrdersRequest) (*ListOpenOrdersResponse, error)
	ModifyOrder(ctx context.Context, request *ModifyOrderRequest) (*ModifyOrderResponse, error)
}

func NewOrdersService(c client.RestClient) OrdersService {
	return &ordersServiceImpl{client: c}
}

type ordersServiceImpl struct {
	client client.RestClient
}
