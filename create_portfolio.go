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
	"github.com/coinbase-samples/core-go"
)

type CreatePortfolioRequest struct {
	Name string `json:"name"`
}

type CreatePortfolioResponse struct {
	Portfolio *Portfolio              `json:"portfolio"`
	Request   *CreatePortfolioRequest `json:"request"`
}

func (c *Client) CreatePortfolio(
	ctx context.Context,
	request *CreatePortfolioRequest,
) (*CreatePortfolioResponse, error) {

	path := "/portfolios"

	response := &CreatePortfolioResponse{Request: request}

	if err := core.Post(ctx, c, path, core.EmptyQueryParams, request, &response.Portfolio, addIntxHeaders); err != nil {
		return nil, err
	}

	return response, nil
}
