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
	"fmt"
)

type ModifyOrderRequest struct {
	OrderId       string `json:"id"`
	PortfolioId   string `json:"portfolio"`
	ClientOrderId string `json:"client_order_id"`
	Price         string `json:"price,omitempty"`
	StopPrice     string `json:"stop_price,omitempty"`
	Size          string `json:"size,omitempty"`
}

type ModifyOrderResponse struct {
	Order   *Order
	Request *ModifyOrderRequest
}

func (c Client) ModifyOrder(
	ctx context.Context,
	request *ModifyOrderRequest,
) (*ModifyOrderResponse, error) {

	path := fmt.Sprintf("/orders/%s", request.OrderId)

	type modifyOrderBody struct {
		Portfolio     string `json:"portfolio"`
		ClientOrderId string `json:"client_order_id"`
		Price         string `json:"price,omitempty"`
		StopPrice     string `json:"stop_price,omitempty"`
		Size          string `json:"size,omitempty"`
	}

	body := modifyOrderBody{
		Portfolio:     request.PortfolioId,
		ClientOrderId: request.ClientOrderId,
		Price:         request.Price,
		StopPrice:     request.StopPrice,
		Size:          request.Size,
	}

	response := &ModifyOrderResponse{Request: request}

	if err := put(ctx, c, path, emptyQueryParams, body, &response.Order); err != nil {
		return nil, err
	}

	return response, nil
}
