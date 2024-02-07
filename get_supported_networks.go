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
)

type GetSupportedNetworksRequest struct {
	AssetId string `json:"asset"`
}

type GetSupportedNetworksResponse struct {
	NetworkDetail []Network                    `json:"network_detail"`
	Request       *GetSupportedNetworksRequest `json:"request"`
}

func (c Client) GetSupportedNetworks(
	ctx context.Context,
	request *GetSupportedNetworksRequest,
) (*GetSupportedNetworksResponse, error) {

	path := fmt.Sprintf("/assets/%s/networks", request.AssetId)

	response := &GetSupportedNetworksResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, nil, &response.NetworkDetail); err != nil {
		return nil, err
	}

	return response, nil
}
