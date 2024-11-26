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

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/intx-sdk-go/client"
	"github.com/coinbase-samples/intx-sdk-go/model"
	"github.com/coinbase-samples/intx-sdk-go/utils"
)

type ListOpenOrdersRequest struct {
	PortfolioId   string `json:"portfolio"`
	InstrumentId  string `json:"instrument,omitempty"`
	ClientOrderId string `json:"client_order_id,omitempty"`
	EventType     string `json:"event_type,omitempty"`
	RefDatetime   string `json:"ref_datetime,omitempty"`
	ResultLimit   int    `json:"result_limit,omitempty"`
	ResultOffset  int    `json:"result_offset,omitempty"`
	Pagination    *model.PaginationParams
}

type ListOpenOrdersResponse struct {
	Pagination *model.PaginationParams `json:"pagination"`
	Results    []*model.Order          `json:"results"`
	Request    *ListOpenOrdersRequest  `json:"request"`
}

func (s ordersServiceImpl) ListOpenOrders(
	ctx context.Context,
	request *ListOpenOrdersRequest,
) (*ListOpenOrdersResponse, error) {

	path := "/orders"

	queryParams := core.AppendHttpQueryParam("", "portfolio", request.PortfolioId)

	if request.InstrumentId != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "instrument", request.InstrumentId)
	}
	if request.ClientOrderId != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "client_order_id", request.ClientOrderId)
	}
	if request.EventType != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "event_type", request.EventType)
	}
	if request.RefDatetime != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "ref_datetime", request.RefDatetime)
	}
	if request.ResultLimit != 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "result_limit", fmt.Sprint(request.ResultLimit))
	}
	if request.ResultOffset != 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "result_offset", fmt.Sprint(request.ResultOffset))
	}

	queryParams = utils.AppendPaginationParams(queryParams, request.Pagination)

	response := &ListOpenOrdersResponse{Request: request}

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		queryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		response,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
