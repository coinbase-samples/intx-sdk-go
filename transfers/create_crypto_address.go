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

type CreateCryptoAddressRequest struct {
	// Deprecated: Use Portfolio instead.
	PortfolioId string `json:"portfolio"`
	Portfolio   string `json:"portfolio"`
	// Deprecated: Use Asset instead.
	AssetId      string `json:"asset"`
	Asset        string `json:"asset"`
	NetworkArnId string `json:"network_arn_id"`
}

type CreateCryptoAddressResponse struct {
	Address *model.Address              `json:"address"`
	Request *CreateCryptoAddressRequest `json:"request"`
}

func (s transfersServiceImpl) CreateCryptoAddress(
	ctx context.Context,
	request *CreateCryptoAddressRequest,
) (*CreateCryptoAddressResponse, error) {

	if request.Portfolio == "" && request.PortfolioId != "" {
		request.Portfolio = request.PortfolioId
	}
	request.PortfolioId = ""

	if request.Asset == "" && request.AssetId != "" {
		request.Asset = request.AssetId
	}
	request.AssetId = ""

	path := "/transfers/address"

	response := &CreateCryptoAddressResponse{Request: request}

	if err := core.HttpPost(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		response,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
