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

type GetPortfolioBalancesRequest struct {
	PortfolioId string `json:"portfolio"`
}

type GetPortfolioBalancesResponse struct {
	Balances []Balance                    `json:"balances"`
	Request  *GetPortfolioBalancesRequest `json:"request"`
}

func (c Client) GetPortfolioBalances(
	ctx context.Context,
	request *GetPortfolioBalancesRequest,
) (*GetPortfolioBalancesResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/balances", request.PortfolioId)

	response := &GetPortfolioBalancesResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, nil, &response.Balances); err != nil {
		return nil, err
	}

	return response, nil
}
