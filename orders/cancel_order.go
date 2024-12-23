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
	"fmt"
	"github.com/coinbase-samples/intx-sdk-go/utils"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/intx-sdk-go/client"
	"github.com/coinbase-samples/intx-sdk-go/model"
)

type CancelOrderRequest struct {
	PortfolioId string `json:"portfolio"`
	// Deprecated: Use Id instead.
	OrderId string `json:"-"`
	Id      string `json:"id"`
}

type CancelOrderResponse struct {
	Order   *model.Order
	Request *CancelOrderRequest
}

func (s ordersServiceImpl) CancelOrder(
	ctx context.Context,
	request *CancelOrderRequest,
) (*CancelOrderResponse, error) {

	utils.FallbackDeprecatedField(&request.Id, request.OrderId)

	path := fmt.Sprintf("/orders/%s", request.Id)

	queryParams := core.AppendHttpQueryParam("", "portfolio", request.PortfolioId)

	response := &CancelOrderResponse{Request: request}

	if err := core.HttpDelete(
		ctx,
		s.client,
		path,
		queryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&response.Order,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
