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

type CreatePortfolioTransferRequest struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Asset  string `json:"asset"`
	Amount string `json:"amount"`
}

type CreatePortfolioTransferResponse struct {
	Success bool                            `json:"success"`
	Request *CreatePortfolioTransferRequest `json:"request"`
}

func (c Client) CreatePortfolioTransfer(
	ctx context.Context,
	request *CreatePortfolioTransferRequest,
) (*CreatePortfolioTransferResponse, error) {

	path := "/portfolios/transfer"

	response := &CreatePortfolioTransferResponse{Request: request}

	if err := post(ctx, c, path, emptyQueryParams, request, &response); err != nil {
		return nil, err
	}

	return response, nil
}
