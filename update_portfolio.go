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

type UpdatePortfolioRequest struct {
	Name        string `json:"name"`
	PortfolioId string `json:"portfolio"`
}

type UpdatePortfolioResponse struct {
	Portfolio *Portfolio              `json:"portfolio"`
	Request   *UpdatePortfolioRequest `json:"request"`
}

func (c *Client) UpdatePortfolio(
	ctx context.Context,
	request *UpdatePortfolioRequest,
) (*UpdatePortfolioResponse, error) {

	path := fmt.Sprintf("/portfolios/%s", request.PortfolioId)

	response := &UpdatePortfolioResponse{Request: request}

	if err := core.Put(ctx, c, path, core.EmptyQueryParams, request, &response.Portfolio, addIntxHeaders); err != nil {
		return nil, err
	}

	return response, nil
}
