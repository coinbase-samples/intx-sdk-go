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

type GetPortfolioPositionsRequest struct {
	PortfolioId string `json:"portfolio"`
}

type GetPortfolioPositionsResponse struct {
	Positions []Position                    `json:"positions"`
	Request   *GetPortfolioPositionsRequest `json:"request"`
}

func (c *Client) GetPortfolioPositions(
	ctx context.Context,
	request *GetPortfolioPositionsRequest,
) (*GetPortfolioPositionsResponse, error) {

	path := fmt.Sprintf("/portfolios/%s/positions", request.PortfolioId)

	response := &GetPortfolioPositionsResponse{Request: request}

	if err := core.Get(ctx, c, path, core.EmptyQueryParams, nil, &response.Positions, addIntxHeaders); err != nil {
		return nil, err
	}

	return response, nil
}
