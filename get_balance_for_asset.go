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

type GetAssetBalanceRequest struct {
	PortfolioId string `json:"portfolio"`
	AssetId     string `json:"asset"`
}

type GetAssetBalanceResponse struct {
	Balance *Balance                `json:"balance"`
	Request *GetAssetBalanceRequest `json:"request"`
}

func (c *Client) GetAssetBalance(
	ctx context.Context,
	request *GetAssetBalanceRequest,
) (*GetAssetBalanceResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/balances/%s", request.PortfolioId, request.AssetId)

	response := &GetAssetBalanceResponse{Request: request}

	if err := core.Get(ctx, c, path, core.EmptyQueryParams, nil, &response.Balance, addIntxHeaders); err != nil {
		return nil, err
	}

	return response, nil
}
