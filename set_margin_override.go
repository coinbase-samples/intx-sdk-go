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
	"errors"
)

type SetMarginOverrideRequest struct {
	Portfolio      string `json:"portfolio"`
	MarginOverride string `json:"margin_override"`
}

type SetMarginOverrideResponse struct {
	MarginOverride *MarginOverride           `json:"margin_override"`
	Request        *SetMarginOverrideRequest `json:"request"`
}

func (c Client) SetMarginOverride(ctx context.Context,
	request *SetMarginOverrideRequest) (*SetMarginOverrideResponse, error) {
	if request == nil {
		return nil, errors.New("set margin override is nil")
	}

	path := "/portfolios/margin"

	response := &SetMarginOverrideResponse{Request: request}

	if err := post(ctx, c, path, emptyQueryParams, request, response); err != nil {
		return nil, err
	}

	return response, nil
}
