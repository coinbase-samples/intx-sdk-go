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

package transfers

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/intx-sdk-go/client"
	"github.com/coinbase-samples/intx-sdk-go/model"
	"github.com/coinbase-samples/intx-sdk-go/utils"
)

type ListTransfersRequest struct {
	// Deprecated: Use Portfolios instead.
	PortfolioIds string `json:"-"`
	Portfolios   string `json:"portfolios"`
	TimeFrom     string `json:"time_from"`
	TimeTo       string `json:"time_to"`
	ResultLimit  int    `json:"result_limit"`
	ResultOffset int    `json:"result_offset"`
	Status       string `json:"status"`
	Type         string `json:"type"`
	Pagination   *model.PaginationParams
}

type ListTransfersResponse struct {
	Pagination *model.PaginationSubset `json:"pagination"`
	Transfers  []*model.Transfer       `json:"results"`
	Request    *ListTransfersRequest   `json:"request"`
}

func (s transfersServiceImpl) ListTransfers(
	ctx context.Context,
	request *ListTransfersRequest,
) (*ListTransfersResponse, error) {

	utils.FallbackDeprecatedField(&request.Portfolios, request.PortfolioIds)

	path := "/transfers"

	queryParams := core.AppendHttpQueryParam("", "portfolios", request.Portfolios)

	if request.TimeFrom != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "time_from", request.TimeFrom)
	}
	if request.TimeTo != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "time_to", request.TimeTo)
	}
	if request.ResultLimit != 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "result_limit", fmt.Sprint(request.ResultLimit))
	}
	if request.ResultOffset != 0 {
		queryParams = core.AppendHttpQueryParam(queryParams, "result_offset", fmt.Sprint(request.ResultOffset))
	}
	if request.Status != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "status", request.Status)
	}
	if request.Type != "" {
		queryParams = core.AppendHttpQueryParam(queryParams, "type", request.Type)
	}

	queryParams = utils.AppendPaginationParams(queryParams, request.Pagination)

	response := &ListTransfersResponse{Request: request}

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		queryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&response.Transfers,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
