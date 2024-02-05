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

type GetHistoricalFundingRequest struct {
	Instrument string `json:"instrument"`
}

type GetHistoricalFundingResponse struct {
	HistoricalFundingRates *HistoricalFundingRate       `json:"historical_funding_rates"`
	Request                *GetHistoricalFundingRequest `json:"request"`
}

func (c Client) GetHistoricalFundingRates(
	ctx context.Context,
	request *GetHistoricalFundingRequest,
) (*GetHistoricalFundingResponse, error) {

	path := fmt.Sprintf("/instruments/%s/funding", request.Instrument)

	response := &GetHistoricalFundingResponse{Request: request}

	if err := get(ctx, c, path, emptyQueryParams, nil, &response.HistoricalFundingRates); err != nil {
		return nil, err
	}

	return response, nil
}
