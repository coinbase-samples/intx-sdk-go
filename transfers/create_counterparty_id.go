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

package transfers

import (
	"context"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/intx-sdk-go/client"
	"github.com/coinbase-samples/intx-sdk-go/model"
)

type CreateCounterpartyIdRequest struct {
	PortfolioId string `json:"portfolio"`
}

type CreateCounterpartyIdResponse struct {
	Counterparty *model.Counterparty          `json:"counterparty"`
	Request      *CreateCounterpartyIdRequest `json:"request"`
}

func (s transfersServiceImpl) CreateCounterpartyId(
	ctx context.Context,
	request *CreateCounterpartyIdRequest,
) (*CreateCounterpartyIdResponse, error) {

	path := "/transfers/create-counterparty-id"

	queryParams := core.AppendHttpQueryParam("", "portfolio", request.PortfolioId)

	response := &CreateCounterpartyIdResponse{Request: request}

	if err := core.HttpPost(
		ctx,
		s.client,
		path,
		queryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		response,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
