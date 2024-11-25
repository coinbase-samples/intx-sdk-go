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
)

type GetPortfolioBalancesRequest struct {
	PortfolioId string `json:"portfolio"`
}

type GetPortfolioBalancesResponse struct {
	Balances []*model.Balance             `json:"balances"`
	Request  *GetPortfolioBalancesRequest `json:"request"`
}

func (s portfoliosServiceImpl) GetPortfolioBalances(
	ctx context.Context,
	request *GetPortfolioBalancesRequest,
) (*GetPortfolioBalancesResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/balances", request.PortfolioId)

	response := &GetPortfolioBalancesResponse{Request: request, Balances: make([]*model.Balance, 0)}

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		response.Balances,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
