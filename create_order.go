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

package intx

import (
	"context"
)

type CreateOrderRequest struct {
	ClientOrderId string  `json:"client_order_id"`
	Side          string  `json:"side"`
	Size          string  `json:"size"`
	Tif           string  `json:"tif"`
	Instrument    string  `json:"instrument"`
	Type          string  `json:"type"`
	Price         string  `json:"price,omitempty"`
	StopPrice     *string `json:"stop_price,omitempty"`
	ExpireTime    *string `json:"expire_time,omitempty"`
	Portfolio     string  `json:"portfolio"`
	User          *string `json:"user,omitempty"`
	StpMode       *string `json:"stp_mode,omitempty"`
	PostOnly      *bool   `json:"post_only,omitempty"`
}

type CreateOrderResponse struct {
	Order   *Order              `json:"order"`
	Request *CreateOrderRequest `json:"request"`
}

func (c Client) CreateOrder(ctx context.Context,
	request *CreateOrderRequest) (*CreateOrderResponse, error) {

	path := "/orders"

	response := &CreateOrderResponse{Request: request}

	if err := post(ctx, c, path, emptyQueryParams, request, &response.Order); err != nil {
		return nil, err
	}

	return response, nil
}
