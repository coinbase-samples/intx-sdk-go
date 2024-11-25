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

package assets

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/intx-sdk-go/client"
	"github.com/coinbase-samples/intx-sdk-go/model"
)

type GetSupportedNetworksRequest struct {
	AssetId string `json:"asset"`
}

type GetSupportedNetworksResponse struct {
	NetworkDetail []*model.Network             `json:"network_detail"`
	Request       *GetSupportedNetworksRequest `json:"request"`
}

func (s assetsServiceImpl) GetSupportedNetworks(
	ctx context.Context,
	request *GetSupportedNetworksRequest,
) (*GetSupportedNetworksResponse, error) {

	path := fmt.Sprintf("/assets/%s/networks", request.AssetId)

	response := &GetSupportedNetworksResponse{Request: request, NetworkDetail: make([]*model.Network, 0)}

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		response.NetworkDetail,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
