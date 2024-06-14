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
	"github.com/coinbase-samples/core-go"
)

type CancelOrderRequest struct {
	PortfolioId string `json:"portfolio"`
	OrderId     string `json:"id"`
}

type CancelOrderResponse struct {
	Order   *Order
	Request *CancelOrderRequest
}

func (c *Client) CancelOrder(
	ctx context.Context,
	request *CancelOrderRequest,
) (*CancelOrderResponse, error) {

	path := fmt.Sprintf("/orders/%s", request.OrderId)

	queryParams := appendQueryParam("", "portfolio", request.PortfolioId)

	response := &CancelOrderResponse{Request: request}

	if err := core.Delete(ctx, c, path, queryParams, nil, &response.Order, addIntxHeaders); err != nil {
		return nil, err
	}

	return response, nil
}
