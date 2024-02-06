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

type CreateWithdrawalToCryptoAddressRequest struct {
	PortfolioId          string `json:"portfolio"`
	AssetId              string `json:"asset"`
	Amount               string `json:"amount"`
	AddNetworkFeeToTotal bool   `json:"add_network_fee_to_total"`
	NetworkArnId         string `json:"network_arn_id"`
	Address              string `json:"address"`
	Nonce                string `json:"nonce"`
}

type CreateWithdrawalToCryptoAddressResponse struct {
	Idem    string                                  `json:"idem"`
	Request *CreateWithdrawalToCryptoAddressRequest `json:"request"`
}

func (c Client) CreateWithdrawalToCryptoAddress(
	ctx context.Context,
	request *CreateWithdrawalToCryptoAddressRequest,
) (*CreateWithdrawalToCryptoAddressResponse, error) {

	path := "/transfers/withdraw"

	response := &CreateWithdrawalToCryptoAddressResponse{Request: request}

	if err := post(ctx, c, path, emptyQueryParams, request, &response); err != nil {
		return nil, err
	}

	return response, nil
}
