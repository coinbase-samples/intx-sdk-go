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

type ListOpenOrdersRequest struct {
	Portfolio     string `json:"portfolio"`
	Instrument    string `json:"instrument,omitempty"`
	ClientOrderId string `json:"client_order_id,omitempty"`
	EventType     string `json:"event_type,omitempty"`
	RefDatetime   string `json:"ref_datetime,omitempty"`
	ResultLimit   int    `json:"result_limit,omitempty"`
	ResultOffset  int    `json:"result_offset,omitempty"`
	Pagination    *PaginationParams
}

type ListOpenOrdersResponse struct {
	Pagination PaginationParams       `json:"pagination"`
	Results    []Order                `json:"results"`
	Request    *ListOpenOrdersRequest `json:"request"`
}

func (c Client) ListOpenOrders(
	ctx context.Context,
	request *ListOpenOrdersRequest,
) (*ListOpenOrdersResponse, error) {

	path := "/orders"

	queryParams := appendQueryParam("", "portfolio", request.Portfolio)

	if request.Instrument != "" {
		queryParams = appendQueryParam(queryParams, "instrument", request.Instrument)
	}
	if request.ClientOrderId != "" {
		queryParams = appendQueryParam(queryParams, "client_order_id", request.ClientOrderId)
	}
	if request.EventType != "" {
		queryParams = appendQueryParam(queryParams, "event_type", request.EventType)
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

	queryParams = appendPaginationParams(queryParams, request.Pagination)

	response := &ListOpenOrdersResponse{Request: request}

	if err := get(ctx, c, path, queryParams, nil, response); err != nil {
		return nil, err
	}

	return response, nil
}
