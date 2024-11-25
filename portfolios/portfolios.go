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

	"github.com/coinbase-samples/intx-sdk-go/client"
)

type PortfoliosService interface {
	CreatePortfolio(ctx context.Context, request *CreatePortfolioRequest) (*CreatePortfolioResponse, error)
	CreatePortfolioTransfer(ctx context.Context, request *CreatePortfolioTransferRequest) (*CreatePortfolioTransferResponse, error)
	GetAssetBalance(ctx context.Context, request *GetAssetBalanceRequest) (*GetAssetBalanceResponse, error)
	GetPortfolio(ctx context.Context, request *GetPortfolioRequest) (*GetPortfolioResponse, error)
	GetPortfolioBalances(ctx context.Context, request *GetPortfolioBalancesRequest) (*GetPortfolioBalancesResponse, error)
	GetPortfolioDetails(ctx context.Context, request *GetPortfolioDetailsRequest) (*GetPortfolioDetailsResponse, error)
	GetPortfolioFills(ctx context.Context, request *GetPortfolioFillsRequest) (*GetPortfolioFillsResponse, error)
	GetPortfolioPositions(ctx context.Context, request *GetPortfolioPositionsRequest) (*GetPortfolioPositionsResponse, error)
	GetPortfolioSummary(ctx context.Context, request *GetPortfolioSummaryRequest) (*GetPortfolioSummaryResponse, error)
	GetInstrumentPosition(ctx context.Context, request *GetInstrumentPositionRequest) (*GetInstrumentPositionResponse, error)
	ListFillsByPortfolios(ctx context.Context, request *ListFillsByPortfoliosRequest) (*ListFillsByPortfoliosResponse, error)
	ListPortfolios(ctx context.Context, request *ListPortfoliosRequest) (*ListPortfoliosResponse, error)
	SetMarginOverride(ctx context.Context, request *SetMarginOverrideRequest) (*SetMarginOverrideResponse, error)
	UpdatePortfolio(ctx context.Context, request *UpdatePortfolioRequest) (*UpdatePortfolioResponse, error)
}

func NewPortfoliosService(c client.RestClient) PortfoliosService {
	return &portfoliosServiceImpl{client: c}
}

type portfoliosServiceImpl struct {
	client client.RestClient
}
