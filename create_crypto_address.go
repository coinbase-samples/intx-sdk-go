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

type CreateCryptoAddressRequest struct {
	Portfolio    string `json:"portfolio"`
	Asset        string `json:"asset"`
	NetworkArnId string `json:"network_arn_id"`
}

type CreateCryptoAddressResponse struct {
	Address *Address                    `json:"address"`
	Request *CreateCryptoAddressRequest `json:"request"`
}

func (c Client) CreateCryptoAddress(
	ctx context.Context,
	request *CreateCryptoAddressRequest,
) (*CreateCryptoAddressResponse, error) {

	path := "/transfers/address"

	response := &CreateCryptoAddressResponse{Request: request}

	if err := post(ctx, c, path, emptyQueryParams, request, &response); err != nil {
		return nil, err
	}

	return response, nil
}
