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

type CancelOrdersRequest struct {
	Portfolio  string `json:"portfolio"`
	Instrument string `json:"instrument"`
}

type CancelOrdersResponse struct {
	Order   []Order
	Request *CancelOrdersRequest
}

func (c Client) CancelOrders(ctx context.Context,
	request *CancelOrdersRequest) (*CancelOrdersResponse, error) {

	path := "/orders"

	queryParams := appendQueryParam("", "portfolio", request.Portfolio)

	if request.Instrument != "" {
		queryParams = appendQueryParam(queryParams, "instrument", request.Instrument)
	}

	response := &CancelOrdersResponse{Request: request}

	if err := del(ctx, c, path, queryParams, nil, &response.Order); err != nil {
		return nil, err
	}

	return response, nil
}
