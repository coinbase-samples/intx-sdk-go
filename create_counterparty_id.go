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

type CreateCounterpartyIdRequest struct {
	PortfolioId string `json:"portfolio"`
}

type CreateCounterpartyIdResponse struct {
	Counterparty *Counterparty                `json:"counterparty"`
	Request      *CreateCounterpartyIdRequest `json:"request"`
}

func (c *Client) CreateCounterpartyId(
	ctx context.Context,
	request *CreateCounterpartyIdRequest,
) (*CreateCounterpartyIdResponse, error) {

	path := "/transfers/create-counterparty-id"

	queryParams := appendQueryParam("", "portfolio", request.PortfolioId)

	response := &CreateCounterpartyIdResponse{Request: request}

	if err := core.Post(ctx, c, path, queryParams, request, &response, addIntxHeaders); err != nil {
		return nil, err
	}

	return response, nil
}
