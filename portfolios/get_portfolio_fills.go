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

package portfolios

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/intx-sdk-go/client"
	"github.com/coinbase-samples/intx-sdk-go/model"
	"github.com/coinbase-samples/intx-sdk-go/utils"
)

type GetPortfolioFillsRequest struct {
	// Deprecated: Use Portfolio instead.
	PortfolioId   string `json:"-"`
	Portfolio     string `json:"portfolio"`
	OrderId       string `json:"order_id,omitempty"`
	ClientOrderId string `json:"client_order_id,omitempty"`
	RefDatetime   string `json:"ref_datetime,omitempty"`
	ResultLimit   int    `json:"result_limit,omitempty"`
	ResultOffset  int    `json:"result_offset,omitempty"`
	TimeFrom      string `json:"time_from,omitempty"`
	Pagination    *model.PaginationParams
}

type GetPortfolioFillsResponse struct {
	Pagination *model.PaginationParams   `json:"pagination"`
	Results    []*model.Fill             `json:"results"`
	Request    *GetPortfolioFillsRequest `json:"request"`
}

func (s portfoliosServiceImpl) GetPortfolioFills(
	ctx context.Context,
	request *GetPortfolioFillsRequest,
) (*GetPortfolioFillsResponse, error) {

	utils.FallbackDeprecatedField(&request.Portfolio, request.PortfolioId)

	path := fmt.Sprintf("/portfolios/%s/fills", request.Portfolio)

	queryParams := core.AppendHttpQueryParam("", "portfolios", request.PortfolioId)

	if request.OrderId != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "order_id", request.OrderId)
	}
	if request.ClientOrderId != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "client_order_id", request.ClientOrderId)
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
	if request.TimeFrom != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "time_from", fmt.Sprint(request.TimeFrom))
	}

	queryParams = utils.AppendPaginationParams(queryParams, request.Pagination)

	response := &GetPortfolioFillsResponse{Request: request}

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
