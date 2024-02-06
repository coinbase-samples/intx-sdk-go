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

type CreateWithdrawalToCounterpartyIdRequest struct {
	PortfolioId    string `json:"portfolio"`
	CounterpartyId string `json:"counterparty_id"`
	AssetId        string `json:"asset"`
	Amount         string `json:"amount"`
	Nonce          string `json:"nonce"`
}

type CreateWithdrawalToCounterpartyIdResponse struct {
	Withdrawal *CounterpartyWithdrawal                  `json:"withdrawal"`
	Request    *CreateWithdrawalToCounterpartyIdRequest `json:"request"`
}

func (c Client) CreateWithdrawalToCounterpartyId(
	ctx context.Context,
	request *CreateWithdrawalToCounterpartyIdRequest,
) (*CreateWithdrawalToCounterpartyIdResponse, error) {

	path := "/transfers/withdraw/counterparty"

	response := &CreateWithdrawalToCounterpartyIdResponse{Request: request}

	if err := post(ctx, c, path, emptyQueryParams, request, &response); err != nil {
		return nil, err
	}

	return response, nil
}
