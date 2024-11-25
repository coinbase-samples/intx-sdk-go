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

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/intx-sdk-go/client"
	"github.com/coinbase-samples/intx-sdk-go/model"
)

type CancelOrdersRequest struct {
	PortfolioId  string `json:"portfolio"`
	InstrumentId string `json:"instrument"`
}

type CancelOrdersResponse struct {
	Orders  []*model.Order
	Request *CancelOrdersRequest
}

func (s ordersServiceImpl) CancelOrders(
	ctx context.Context,
	request *CancelOrdersRequest,
) (*CancelOrdersResponse, error) {

	path := "/orders"

	queryParams := core.AppendHttpQueryParam("", "portfolio", request.PortfolioId)

	if request.InstrumentId != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "instrument", request.InstrumentId)
	}

	response := &CancelOrdersResponse{Request: request, Orders: make([]*model.Order, 0)}

	if err := core.HttpDelete(
		ctx,
		s.client,
		path,
		queryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		response.Orders,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
