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

type GetPortfolioFillsRequest struct {
	PortfolioId   string `json:"portfolio"`
	OrderId       string `json:"order_id,omitempty"`
	ClientOrderId string `json:"client_order_id,omitempty"`
	RefDatetime   string `json:"ref_datetime,omitempty"`
	ResultLimit   int    `json:"result_limit,omitempty"`
	ResultOffset  int    `json:"result_offset,omitempty"`
	TimeFrom      string `json:"time_from,omitempty"`
	Pagination    *PaginationParams
}

type GetPortfolioFillsResponse struct {
	Pagination PaginationParams          `json:"pagination"`
	Results    []Fill                    `json:"results"`
	Request    *GetPortfolioFillsRequest `json:"request"`
}

func (c *Client) GetPortfolioFills(
	ctx context.Context,
	request *GetPortfolioFillsRequest,
) (*GetPortfolioFillsResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/fills", request.PortfolioId)

	queryParams := appendQueryParam("", "portfolios", request.PortfolioId)

	if request.OrderId != "" {
		queryParams = appendQueryParam(queryParams, "order_id", request.OrderId)
	}
	if request.ClientOrderId != "" {
		queryParams = appendQueryParam(queryParams, "client_order_id", request.ClientOrderId)
	}
	if request.RefDatetime != "" {
		queryParams = appendQueryParam(queryParams, "ref_datetime", request.RefDatetime)
	}
	if request.ResultLimit != 0 {
		queryParams = appendQueryParam(queryParams, "result_limit", fmt.Sprint(request.ResultLimit))
	}
	if request.ResultOffset != 0 {
		queryParams = appendQueryParam(queryParams, "result_offset", fmt.Sprint(request.ResultOffset))
	}
	if request.TimeFrom != "" {
		queryParams = appendQueryParam(queryParams, "time_from", fmt.Sprint(request.TimeFrom))
	}

	queryParams = appendPaginationParams(queryParams, request.Pagination)

	response := &GetPortfolioFillsResponse{Request: request}

	if err := core.Get(ctx, c, path, queryParams, nil, response, addIntxHeaders); err != nil {
		return nil, err
	}

	return response, nil
}
