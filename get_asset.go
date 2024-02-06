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

type GetAssetRequest struct {
	AssetId string `json:"asset"`
}

type GetAssetResponse struct {
	AssetDetail *Asset
	Request     *GetAssetRequest
}

func (c Client) GetAsset(
	ctx context.Context,
	request *GetAssetRequest,
) (*GetAssetResponse, error) {

	path := fmt.Sprintf("/assets/%s", request.AssetId)

	response := &GetAssetResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, nil, &response.AssetDetail); err != nil {
		return nil, err
	}

	return response, nil
}
