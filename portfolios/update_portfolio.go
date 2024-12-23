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

package portfolios

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/intx-sdk-go/client"
	"github.com/coinbase-samples/intx-sdk-go/model"
)

type UpdatePortfolioRequest struct {
	Name string `json:"name"`
	// Deprecated: Use Portfolio instead.
	PortfolioId string `json:"portfolio"`
	Portfolio   string `json:"portfolio"`
}

type UpdatePortfolioResponse struct {
	Portfolio *model.Portfolio        `json:"portfolio"`
	Request   *UpdatePortfolioRequest `json:"request"`
}

func (s portfoliosServiceImpl) UpdatePortfolio(
	ctx context.Context,
	request *UpdatePortfolioRequest,
) (*UpdatePortfolioResponse, error) {

	portfolio := request.Portfolio
	if portfolio == "" {
		portfolio = request.PortfolioId
	}

	path := fmt.Sprintf("/portfolios/%s", portfolio)

	response := &UpdatePortfolioResponse{Request: request}

	if err := core.HttpPost(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&response.Portfolio,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
