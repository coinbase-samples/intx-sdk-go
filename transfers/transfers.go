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

package transfers

import (
	"context"

	"github.com/coinbase-samples/intx-sdk-go/client"
)

type TransfersService interface {
	CreateCounterpartyId(ctx context.Context, request *CreateCounterpartyIdRequest) (*CreateCounterpartyIdResponse, error)
	CreateCryptoAddress(ctx context.Context, request *CreateCryptoAddressRequest) (*CreateCryptoAddressResponse, error)
	CreateWithdrawalToCounterpartyId(ctx context.Context, request *CreateWithdrawalToCounterpartyIdRequest) (*CreateWithdrawalToCounterpartyIdResponse, error)
	CreateWithdrawalToCryptoAddress(ctx context.Context, request *CreateWithdrawalToCryptoAddressRequest) (*CreateWithdrawalToCryptoAddressResponse, error)
	GetTransfer(ctx context.Context, request *GetTransferRequest) (*GetTransferResponse, error)
	ListTransfers(ctx context.Context, request *ListTransfersRequest) (*ListTransfersResponse, error)
	ValidateCounterpartyId(ctx context.Context, request *ValidateCounterpartyIdRequest) (*ValidateCounterpartyIdResponse, error)
}

func NewTransfersService(c client.RestClient) TransfersService {
	return &transfersServiceImpl{client: c}
}

type transfersServiceImpl struct {
	client client.RestClient
}
