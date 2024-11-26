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

package instruments

import (
	"context"
	"fmt"

	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/intx-sdk-go/client"
	"github.com/coinbase-samples/intx-sdk-go/model"
)

type GetInstrumentQuoteRequest struct {
	InstrumentId string `json:"instrument"`
}

type GetInstrumentQuoteResponse struct {
	InstrumentQuote *model.Quote               `json:"instrument_quote"`
	Request         *GetInstrumentQuoteRequest `json:"request"`
}

func (s instrumentsServiceImpl) GetInstrumentQuote(
	ctx context.Context,
	request *GetInstrumentQuoteRequest,
) (*GetInstrumentQuoteResponse, error) {

	path := fmt.Sprintf("/instruments/%s/quote", request.InstrumentId)

	response := &GetInstrumentQuoteResponse{Request: request}

	if err := core.HttpGet(
		ctx,
		s.client,
		path,
		core.EmptyQueryParams,
		client.DefaultSuccessHttpStatusCodes,
		request,
		&response.InstrumentQuote,
		s.client.HeadersFunc(),
	); err != nil {
		return nil, err
	}

	return response, nil
}
